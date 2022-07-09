package main

import "fmt"

type Currency int

const (
	USD Currency = iota
	EUR
	GBP
	RMB
)

func main() {
	symbol := [...]string{EUR: "€", GBP: "£", RMB: "¥"}
	fmt.Println(len(symbol), symbol)
}
