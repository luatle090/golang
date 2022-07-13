package main

import "fmt"

func main() {

	res := make(chan int)

	go func() {
		//time.Sleep(time.Second * 10)

		res <- 12
	}()

	go func() {
		//time.Sleep(time.Second * 10)

		res <- 12
	}()

	for ss := range res {
		fmt.Println(ss)
	}

}
