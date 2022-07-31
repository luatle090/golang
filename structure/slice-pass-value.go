package main

import "fmt"

func PassByValue(a int, list []int) {

	// mỗi lần append sẽ trả về 1 con trỏ mới
	for i := a; i > 0; i-- {
		list = append(list, i)
	}
}

func RightThing(a int, list []int) []int {
	for i := a; i > 0; i-- {
		list = append(list, i)
	}
	return list
}

// tìm con của start
func PassByValueRecursion(start int, parents map[int]int, list []int) {
	for key, parent := range parents {
		if start == parent {
			list = append(list, key)
			PassByValueRecursion(key, parents, list)
		}
	}
}

// tìm con của start
func RightThingRecursion(start int, parents map[int]int, list []int) []int {
	for key, parent := range parents {
		if start == parent {
			list = append(list, key)
			list = RightThingRecursion(key, parents, list)
		}
	}
	return list
}

func main() {
	list := make([]int, 0)
	PassByValue(5, list)

	list = RightThing(5, list)

	// for i := range list {
	// 	fmt.Println("element", i, ":", list[i])
	// }

	// ------------------------------------------
	// ------------------------------------------

	// value 0 is parent
	parent := map[int]int{
		1:  0,
		2:  0,
		3:  1,
		4:  1,
		5:  2,
		6:  3,
		7:  3,
		8:  5,
		9:  4,
		10: 7,
	}

	list2 := make([]int, 0)
	PassByValueRecursion(3, parent, list2)
	fmt.Println(len(list2))

	list2 = RightThingRecursion(1, parent, list2)

	fmt.Println(len(list2))
	for _, v := range list2 {
		fmt.Println("con cua", 1, ":", v)
	}
}
