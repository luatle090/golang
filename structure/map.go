package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	// dùng thư viện để slipt string giữa các khoảng trắng
	chuoi := strings.Fields(s)

	// khai báo map
	result := make(map[string]int)
	for _, v := range chuoi {
		if _, isHasKey := result[v]; !isHasKey {
			result[v] = 1
		} else {
			result[v]++
		}
	}
	return result
}

func main() {
	wc.Test(WordCount)
}
