package main

import "fmt"

func main() {

	var a []int = make([]int, 64)
	count := 1
	for y := 0; y < 8; y++ {
		for x := 0; x < 8; x++ {
			fmt.Println(a[y*8+x], count)
			count++
		}
	}
}
