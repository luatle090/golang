package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

type byteCounter struct {
	w     io.Writer
	count int64
}

func main() {
	data := []byte("Ahi there")

	// sử dụng 1 con trỏ buffer để write. Ghi vào memory (hàm write đã implement)
	b := &bytes.Buffer{}

	// 1 implement cụ thể
	n, err := b.Write(data)
	if err != nil {
		return
	}

	fmt.Printf("write %d bytes total %d bytes\n", b, n) // tổng số byte đã ghi

	c, num := CountingWriter(b)
	c.Write(data)

	fmt.Printf("\nwrite %d bytes, total %d bytes\n", c, *num)

	// kiem tra dia chi của interface io.Writer có giống địa chỉ của
	func() {
		fmt.Printf("Dia chi cua bien interface: %p\n", c)
		fmt.Printf("Dia chi gia tri cua bien interface: %+v\n", c)
	}()

	// write vao file
	file, errFile := os.Create("test.txt")
	if errFile != nil {
		fmt.Println(errFile)
	}

	defer file.Close()

	writer, num := CountingWriter(file)

	writer.Write(data)
}

// init value cho newWriter mà wrap writer gốc
// Trả về: Trả ra con trỏ interface trỏ đến struct.
// Nếu *io.Writer nghĩa là con trỏ trỏ đến interface
func CountingWriter(w io.Writer) (io.Writer, *int64) {
	b := byteCounter{w: w, count: 0}
	fmt.Printf("dia chi byteCounter: %p\n", &b)
	return &b, &b.count
}

// implement write của interface writer
// bên trong ko implement cách ghi mà sẽ gọi hàm write mà struct này đâ wrap lại
func (b *byteCounter) Write(p []byte) (int, error) {
	n, err := b.w.Write(p) // gọi hàm write để ghi.
	fmt.Printf("\n (in method) write %d bytes total %d bytes\n", b, n)
	b.count += int64(n)
	return n, err
}
