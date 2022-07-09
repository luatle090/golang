package main

import "fmt"

func fibonaccy(a int) int {
	if a == 1 || a == 2 {
		return 1
	}
	return fibonaccy(a-1) + fibonaccy(a-2)
}

func main() {

	result := fibonaccy(4)
	fmt.Println(result)
}
