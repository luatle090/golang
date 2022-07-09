package main

import (
	"bytes"
	"fmt"
	"io"
	"strings"
)

type NewReader struct {
	r        io.Reader
	n, limit int64
}

func (newReader *NewReader) Read(b []byte) (n int, err error) {
	if x := newReader.limit - newReader.n; int(x) < len(b) {
		n, err = newReader.r.Read(b[:x])
	} else {
		n, err = newReader.r.Read(b)
	}
	n, err = newReader.r.Read(b)
	newReader.n += int64(n)
	if newReader.n > newReader.limit {
		err = io.EOF
	}
	return
}

func LimitReader(r io.Reader, n int64) io.Reader {
	return &NewReader{r: r, limit: n}
}

func main() {
	data := []byte("con cho nuoi con meo")
	b := &bytes.Buffer{}

	// ghi dữ liệu lên memory buffer để có nguồn đọc dữ liệu
	_, err := b.Write(data)

	if err != nil {
		fmt.Println(err)
		return
	}

	buff := make([]byte, 100)

	reader := LimitReader(b, 10)

	for {
		n, errRead := reader.Read(buff)
		fmt.Printf("%d\n", n)
		fmt.Printf("%q\n", buff[:n])

		if errRead == io.EOF {
			break
		}

		if errRead != nil {
			fmt.Println(errRead)
			break
		}

	}
	//Reader()
}

func Reader() {
	r := strings.NewReader("Hello, Reader!")

	b := make([]byte, 8)
	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}
}
