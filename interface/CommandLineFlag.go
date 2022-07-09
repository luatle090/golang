/*

	Hàm mô tả 2 kiểu khai báo để implement flag sử dụng cho tùy biến command line
	Command Line là biến toàn cục (biến toàn cục flag.CommandLine)
	Sách the go programming trang 180
*/

package main

import (
	"flag"
	"fmt"
)

type Celsius float64
type Fahrenheit float64

func FToC(f Fahrenheit) Celsius { return Celsius((f - 32.0) * 5.0 / 9.0) }

/*
// implement flag với 2 phương thức
package flag

// Value is the interface to the value stored in a flag.
type Value interface {
	String() string
	Set(string) error
}
//!-flagvalue
*/

// 		Khai Báo 1:
/*
	Struct celsiusFlag có 1 property là c Celsius
 	vì vậy struct celsiusFlag đều phải implement string và set (sử dụng reciver *celsiusFlag và celsiusFlag) của interface Value
 	Con trỏ trả ra của hàm CelsiusFlag là *celsiusFlag
*/
type celsiusFlag struct{ c Celsius }

func (f *celsiusFlag) Set(s string) error {
	fmt.Println("call")
	var unit string
	var value float64
	fmt.Sscanf(s, "%f%s", &value, &unit) // no error check needed
	switch unit {
	case "C", "°C":
		f.c = Celsius(value)
		return nil
	case "F", "°F":
		f.c = FToC(Fahrenheit(value))
		return nil
	}
	return fmt.Errorf("invalid temperature %q", s)
}

func (c celsiusFlag) String() string { return fmt.Sprintf("%g°C", c.c) }

func CelsiusFlag(name string, value Celsius, usage string) *celsiusFlag {
	f := celsiusFlag{value}
	flag.CommandLine.Var(&f, name, usage)
	return &f
}

// 		Khai báo 2:
/*
	- Sử dụng embedded, struct celsiusFlag khai báo Celsius là 1 embedded field
	- Hàm Set implement như cách khai báo 1. Tuy nhiên khi phép gán f.Celsius thay vì f.c
 	- Hàm String receiver sẽ là Celsius thay vì là celsiusFlag
 	- Hàm CelsiusFlag sẽ trả về *Celsius thay vì là *celsiusFlag như khai báo 1. Vì:
 	- hàm string receiver là Celsius chứ ko phải là celsiusFlag, trả như vậy mới có thể in đc value
*/
// uncomment đoạn dưới để sử dụng
/*
type celsiusFlag struct{ Celsius }

func (f *celsiusFlag) Set(s string) error {
	var unit string
	var value float64
	fmt.Sscanf(s, "%f%s", &value, &unit) // no error check needed
	switch unit {
	case "C", "°C":
		f.Celsius = Celsius(value)
		return nil
	case "F", "°F":
		f.Celsius = FToC(Fahrenheit(value))
		return nil
	}
	return fmt.Errorf("invalid temperature %q", s)
}

// Do tính chất embedded nên Celsius đã đc thăng chức
func (c Celsius) String() string { return fmt.Sprintf("%g°C", c) }

func CelsiusFlag(name string, value Celsius, usage string) *Celsius {
	f := celsiusFlag{value}
	flag.CommandLine.Var(&f, name, usage)
	return &f.Celsius
}
*/

var temp = CelsiusFlag("temp", 20.0, "the temperature")

func main() {
	flag.Parse()
	fmt.Println(*temp)
}
