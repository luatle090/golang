package main

import "fmt"

type Vertex struct {
	x float32
	y float32
}

type PhanSo struct {
	tu  int
	mau int
}

func main() {
	var sum int
	var a, b int

	a = 155
	b = 225
	sum = a + b
	fmt.Println(sum)
	fmt.Println(sumC(a, b))
	fmt.Println(sumLoop())
	exDeferStack()
	Swap(&a, &b)
	fmt.Println("a: ", a, " b: ", b)
	fmt.Println("---- struct trong go lang -----")
	fmt.Println(Vertex{4.5, 6.1})
	v := Vertex{4.5, 6.1}
	fmt.Println(v)
	fmt.Println("x: ", v.x, " y: ", v.y)
	fmt.Println("---- struct - pointer trong go lang -----")

	var ps PhanSo = PhanSo{1, 4}
	// or
	// var ps PhanSo

	fmt.Println("tu so: ", ps.tu, " mau so", ps.mau)
	ps.tu = 1
	ps.mau = 2

	fmt.Println("tu so: ", ps.tu, " mau so", ps.mau)
	pointerPhanSo := &ps
	pointerPhanSo.tu = 2
	pointerPhanSo.mau = 5

	fmt.Println("pointer: tu so: ", ps.tu, " mau so", ps.mau)
}

// return value
func sumC(a int, b int) int {
	return a + b
}

// return value với sum là giá trị duy nhất đc trả về
func sumLoop() (sum int) {
	for i := 0; i < 10; i++ {
		sum += i
	}
	return
}

// will call delay, function call not executed unitl the surrounding fuction return
func exDeferStack() {
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
}

// pointer trong golang, golang ko cho thực hiện phép tính số học như C
func Swap(a *int, b *int) {
	temp := *a
	fmt.Println("temp ", temp)
	*a = *b
	*b = temp
}
