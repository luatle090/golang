package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

type Data struct {
	Offset int
	Key    string
	Value  string
}

const LOG_FILE = "log2.log"

// var vFlag = flag.Bool("set", false, "Set key and value into database")
var nextOffset int = 0
var offsets = make(map[string]Data)

var moveCursorToNext = 1

func main() {
	// flag.Parse()
	// roots := flag.Args()
	// if len(roots) == 0 {
	// 	fmt.Println("input is required")
	// 	return
	// }

	// if *vFlag {
	// 	db_set(roots[0], roots[1])
	// 	return
	// }

	// value, err := db_get(roots[0])
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// fmt.Println(value)

	var input string
	loadIndex()
	for {
		fmt.Println("Press command key value: ")
		scan := bufio.NewScanner(os.Stdin)
		if scan.Scan() {
			input = scan.Text()
		}
		if input == "exit" {
			break
		}

		token := strings.SplitN(input, " ", 3)
		if len(token) < 2 {
			continue
		}

		if token[0] == "set" {
			db_set(token[1], token[2])
			continue
		}

		if token[0] == "get" {
			value, err := db_get(token[1])
			if err != nil {
				log.Fatalln(err)
			}
			fmt.Println(value)
		}

		if token[0] == "merge" {
			if err := db_compaction(); err != nil {
				log.Fatalln(err)
			}
		}
	}
}

// set dữ liệu vào log theo key và value
func db_set(key, value string) error {
	var file *os.File
	var err error
	if !fileExists(LOG_FILE) {
		file, err = os.Create(LOG_FILE)
	} else {
		file, err = os.OpenFile(LOG_FILE, os.O_APPEND|os.O_WRONLY, 0666)
	}

	if err != nil {
		return err
	}
	defer file.Close()

	n, err := fmt.Fprintf(file, "%s,%s\n", key, value)
	hashIndex(key, value, n)
	if err != nil {
		return err
	}
	return nil
}

// lấy dữ liệu từ log theo key
func db_get(key string) (string, error) {
	file, err := os.Open(LOG_FILE)

	if err != nil {
		return "", err
	}

	defer file.Close()

	offset, err := getOffset(key)
	if err != nil {
		return "", err
	}
	// file.Seek(int64(offset), 0)
	// scanner := bufio.NewScanner(file)
	// var valueFromLastKey string

	// if scanner.Scan() {
	// 	str := strings.Split(scanner.Text(), ",")
	// 	valueFromLastKey = str[1]
	// }
	// return valueFromLastKey, nil
	return offset.Value, nil
}

// file có tồn tại
func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

// set key và offset
func hashIndex(key, value string, n int) {
	offsets[key] = Data{Offset: nextOffset, Key: key, Value: value}
	nextOffset += n
	fmt.Println(offsets)
}

// lấy offset từ trong map offset
func getOffset(key string) (Data, error) {
	offset, ok := offsets[key]
	if !ok {
		return Data{}, fmt.Errorf("not found key %s", key)
	}
	return offset, nil
}

// load index có sẵn
func loadIndex() {
	file, err := os.Open(LOG_FILE)
	if err != nil {
		return
	}

	defer file.Close()

	var n int
	scanner := bufio.NewScanner(file)
	// reader := bufio.NewReader(file)
	// for {
	// 	prop, err := reader.ReadBytes('\n')
	// 	if err != nil {
	// 		log.Fatalln(err)
	// 	}
	// 	if err == io.EOF {
	// 		break
	// 	}

	// 	offset := len(prop)
	// 	 string(prop)
	// 	n += offset

	// 	offsets[] = n
	// }
	for scanner.Scan() {
		str := scanner.Text()
		byte := len(str)

		prop := strings.Split(str, ",")
		offsets[prop[0]] = Data{Offset: n, Key: prop[0], Value: prop[1]}
		n += byte + moveCursorToNext

	}
	info, _ := file.Stat()

	nextOffset = int(info.Size())
	fmt.Println()
	fmt.Println(offsets)
}

// TODO: feature Compaction
func db_compaction() error {
	file, err := os.Create(generateFileName())
	if err != nil {
		return err
	}
	for _, data := range offsets {
		_, err := fmt.Fprintf(file, "%s,%s\n", data.Key, data.Value)
		if err != nil {
			return err
		}
	}
	return nil
}

func generateFileName() string {
	current := time.Now()
	// current.Hour()
	// current.Minute()
	// current.Second()
	// current.YearDay()
	// current.Year()
	return fmt.Sprintf("%d-file.log", current.Unix())
}

func (dt Data) String() string {
	return fmt.Sprintf("data %s offset - %d\n", dt.Value, dt.Offset)
}
