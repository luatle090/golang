package main

import "fmt"

func removeIndex(slice []int, s int) []int {
	// lấy vị trí đầu mảng đến vị trí s-1, rồi append với vị trí s+1 đến hết mảng
	return append(slice[:s], slice[s+1:]...)
}

// Remove slice trong vòng for
func main() {
	array := []int{1, 2, 3, 4}

	// for i := range array {
	// 	array = append(array, (i+2+1)*4)
	// 	fmt.Println("index", i)
	// }

	// fmt.Println(array[len(array):])
	// fmt.Println(array[len(array)-1+1:])

	// cả 2 for bên dưới khi remove ko bị crash program
	for i := len(array); i > 0; i-- {
		array = removeIndex(array, len(array)-1)
		fmt.Println(array)
	}

	for range array {
		array = removeIndex(array, len(array)-1)
	}

	// for _, value := range arr {
	// fmt.Println(value)
	// }
}
