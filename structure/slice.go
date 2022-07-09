package main

import (
	"fmt"

	"golang.org/x/tour/pic"
)

// tạo mảng 2 chiều
func Pic(dx, dy int) [][]uint8 {
	arr := make([][]uint8, dy) // tạo arr với chiều dài dy
	for i := range arr {
		arr[i] = make([]uint8, dx) // cấp phát mảng
	}

	for i := range arr {
		for j := range arr[i] {
			arr[i][j] = uint8(i ^ j)
		}
	}

	return arr
}

func ss(b []int) {
	b[0] = 11
}

func main() {
	a := make([]int, 8)
	ss(a)
	for _, v := range a {
		fmt.Println(v)
	}
	pic.Show(Pic)
}
