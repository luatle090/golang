package main

import (
	"flag"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// https://www.w3.org/TR/xml-entity-names/025.html
var Digits = [][]string{
	{" ┏━┓ ", "  ╻  ", " ┏━┓ ", " ━━┓ ", " ╻ ╻ ", " ┏━┓ ", " ┏━┓ ", " ┏━┓ ", " ┏━┓ ", " ┏━┓ ", "   ", "   ", "   "},
	{" ┃ ┃ ", "  ┃  ", "   ┃ ", "   ┃ ", " ┃ ┃ ", " ┃   ", " ┃   ", "   ┃ ", " ┃ ┃ ", " ┃ ┃ ", " ╻ ", "   ", "   "},
	{" ┃ ┃ ", "  ┃  ", "   ┃ ", "   ┃ ", " ┃ ┃ ", " ┃   ", " ┃   ", "   ┃ ", " ┃ ┃ ", " ┃ ┃ ", "   ", "   ", "   "},
	{" ┃ ┃ ", "  ┃  ", " ┏━┛ ", " ━━┫ ", " ┗━┫ ", " ┗━┓ ", " ┣━┓ ", "   ┃ ", " ┣━┫ ", " ┗━┫ ", "   ", " ━ ", "   "},
	{" ┃ ┃ ", "  ┃  ", " ┃   ", "   ┃ ", "   ┃ ", "   ┃ ", " ┃ ┃ ", "   ┃ ", " ┃ ┃ ", "   ┃ ", "   ", "   ", "   "},
	{" ┃ ┃ ", "  ┃  ", " ┃   ", "   ┃ ", "   ┃ ", "   ┃ ", " ┃ ┃ ", "   ┃ ", " ┃ ┃ ", "   ┃ ", " ╹ ", "   ", "   "},
	{" ┗━┛ ", "  ╹  ", " ┗━━ ", " ━━┛ ", "   ╹ ", " ┗━┛ ", " ┗━┛ ", "   ╹ ", " ┗━┛ ", " ┗━┛ ", "   ", "   ", "   "},
}

var digitCheck = regexp.MustCompile(`^[0-9]+$`)
var clock = flag.Bool("date", false, "show date")

func main() {

	// current := time.Now()
	// fmt.Printf("\r%s\n", current.Format("02-01-2006 15:04:05"))
	// t := current.Format("02-01-2006 15:04:05")

	// for _, v := range t {
	// 	if digitCheck.MatchString(string(v)) {
	// 		fmt.Print(v)
	// 	} else {
	// 		fmt.Print("-")
	// 	}
	// }

	// for i := range Digits {
	// 	for j := range `123456789` {
	// 		fmt.Printf("%s", Digits[i][j])
	// 	}
	// 	fmt.Println()
	// }

	flag.Parse()
	fmt.Printf("\x1b[2J")
	fmt.Printf("\x1b[?25l")

	format := "15:04:05"
	if *clock {
		format = "02-01-2006 15:04:05"
	}

	for {
		current := time.Now()
		// fmt.Printf("\r%s", current.Format("02-01-2006 15:04:05"))
		t := current.Format(format)
		for row := range Digits {
			for _, c := range t {
				var col int
				if digitCheck.MatchString(string(c)) {
					col, _ = strconv.Atoi(string(c))
				} else if strings.Compare("-", string(c)) == 0 {
					col = 11
				} else if strings.Compare(":", string(c)) == 0 {
					col = 10
				} else {
					col = 12
				}
				fmt.Printf("%s", Digits[row][col])
			}
			fmt.Println()
		}
		time.Sleep(time.Millisecond * 999)

		fmt.Printf("\x1b[7A")
		// fmt.Printf("\x1b[7A")
		current = current.Add(time.Second * 1)
		//fmt.Println(current.Format(format))
	}
}
