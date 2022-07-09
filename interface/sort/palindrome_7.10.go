package main

import (
	"fmt"
	"sort"
)

func equal(i, j int, s sort.Interface) bool {
	return !s.Less(i, j) && !s.Less(j, i)
}

// kiểm tra đối xứng
func IsPalindrome(s sort.Interface) bool {
	// minus 1 bởi mảng bắt đầu từ 0
	end := s.Len() - 1
	for i := 0; i < s.Len()/2; i++ {
		if !equal(i, end-i, s) {
			return false
		}
	}
	return true
}

func main() {
	v := []int{1, 2, 3, 3, 2, 1}

	fmt.Printf("%v\n", IsPalindrome(sort.IntSlice(v)))

	v = []int{1, 2, 3, 2, 1}
	fmt.Printf("%v\n", IsPalindrome(sort.IntSlice(v)))

	v = []int{1, 1, 2, 3, 2, 1}
	fmt.Printf("%v\n", IsPalindrome(sort.IntSlice(v)))

	v = []int{1, 2, 3, 3, 1, 1}
	fmt.Printf("%v\n", IsPalindrome(sort.IntSlice(v)))

	v = []int{1, 2, 3, 4, 1, 2, 1}
	fmt.Printf("%v\n", IsPalindrome(sort.IntSlice(v)))
}
