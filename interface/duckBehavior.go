package main

import "fmt"

type SizeOfDuck interface {
	len() int
}

type Duck struct {
	name string
}

type Lion struct {
	name string
}

func len(duckSlice []Duck) int {
	return 21
}

func (d Duck) len() int {
	return 21
}

func (l Lion) len() int {
	return 12
}

func behaviorLen() {
	fmt.Println("----- Behavior like len -----")
	duckList := []Duck{{"Duck Name 1"}, {"Duck Name 2"}}

	duckList2 := make([]Duck, 2)

	duckList2[0] = Duck{"Duck name 1 in list 2"}

	fmt.Println(len(duckList))
	fmt.Println(len(duckList2))
	for _, v := range duckList2 {
		fmt.Println(v)
	}
}

func behaviorLikeDuck() {
	fmt.Println("----- Behavior like Duck -----")
	var sizeOfDuck SizeOfDuck
	lion := Lion{"Lion name 1"}
	sizeOfDuck = lion
	fmt.Println(sizeOfDuck.len())
}

func main() {
	behaviorLen()
	behaviorLikeDuck()
}
