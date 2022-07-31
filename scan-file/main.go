package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

var vFlag = flag.Bool("v", false, "show verbose progress messages")

const (
	FILE_NAME   = "file-scan.json"
	FOLDER_NAME = "folder-scan.json"
)

type FileInfor struct {
	Name string
	Path string
}

type FolderInfo struct {
	Name  string
	Count int
}

func walkDir(dir string, folderName string, fileList []FileInfor, folderList []*FolderInfo) ([]FileInfor, []*FolderInfo) {

	folder := FolderInfo{Name: folderName, Count: 0}
	folderList = append(folderList, &folder)

	for _, entry := range dirents(dir) {
		folder.Count++
		if entry.IsDir() {
			subdir := filepath.Join(dir, entry.Name())
			fileList, folderList = walkDir(subdir, entry.Name(), fileList, folderList)
		} else {
			FilePath := filepath.Join(dir, entry.Name())
			scan := FileInfor{Name: entry.Name(), Path: FilePath}
			fileList = append(fileList, scan)
		}
	}
	return fileList, folderList
}

func dirents(dir string) []os.FileInfo {
	entries, err := ioutil.ReadDir(dir)

	if err != nil {
		fmt.Fprintf(os.Stderr, "du: %v\n", err)
		return nil
	}
	return entries
}

func PrintList(fileList []FileInfor) {
	for i := range fileList {
		fmt.Println(fileList[i].Name)
	}

}

func WriteJson[T any](filename string, fileList []T) {
	f, err := os.Create(filename)
	if err != nil {
		log.Println(err)
	}

	//json.Marshal()
	json, err := json.Marshal(fileList)
	if err != nil {
		log.Println(err)
	}

	// 				using streaming
	//// --------------------------------------------------------
	// enc := json.NewEncoder(f)
	// if err := enc.Encode(&fileList); err != nil {
	// 	log.Println(err)
	// }
	//// --------------------------------------------------------

	_, err = f.Write(json)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()
}

func main() {
	flag.Parse()

	// Determine the initial directories.
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}

	fileList := make([]FileInfor, 0)
	var folderList []*FolderInfo
	for _, root := range roots {
		fileList, folderList = walkDir(root, root, fileList, folderList)
	}

	fmt.Println(FOLDER_NAME)

	WriteJson(FILE_NAME, fileList)
	WriteJson(FOLDER_NAME, folderList)
}
