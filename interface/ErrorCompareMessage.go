package main

import (
	"fmt"
	"strconv"
)

func New(text string) error { return errorMessage{text} }

type errorMessage struct{ text string }

func (e errorMessage) Error() string {
	return e.text
}

func NewErrorPointer(text string) error { return &errorMessagePointer{text} }

type errorMessagePointer struct{ text string }

// hàm là pointer receiver
func (e *errorMessagePointer) Error() string {
	return e.text
}

// lý do ErrorString trả về pointer
func main() {
	// ko trả về pointer
	err1 := New("EOF")
	err2 := New("EOF")

	fmt.Println("non pointer: " + strconv.FormatBool(err1 == err2))
	fmt.Printf("non pointer 1: %v , non pointer 2: %v\n", err1, err2)

	fmt.Println("-------------")
	// trả về pointer
	errPointer1 := NewErrorPointer("EOF")
	errPointer2 := NewErrorPointer("EOF")

	fmt.Println("pointer: " + strconv.FormatBool(errPointer1 == errPointer2))

	// lý do vì so sánh địa chỉ đang khác nhau, tuy type là giống nhau
	fmt.Printf("pointer 1: %p , pointer 2: %p\n", errPointer1, errPointer2)
}
