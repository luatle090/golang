package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)      // doc tu stdin
	scanner.Scan()                             // nhan stdin (ban phim). Tra ve true or false neu con nhan tiep
	number, ok := strconv.Atoi(scanner.Text()) // lay text va parse ra kieu int

	if ok != nil {
		fmt.Println(ok)
	}
	number += number
	fmt.Println(number)

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}
