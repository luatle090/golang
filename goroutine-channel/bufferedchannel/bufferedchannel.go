package main

import (
	"fmt"
)

func main() {
	res := make(chan string, 3)
	fmt.Println(mirroredQueue(res))
	for i := 0; i < cap(res)-1; i++ {
		fmt.Println(<-res)
	}

}

func mirroredQueue(res chan string) string {
	go func() { res <- request("asia") }()
	go func() { res <- request("europe") }()
	go func() { res <- request("america") }()

	return <-res
}

func request(hostname string) string {
	switch hostname {
	case "asia":
		//time.Sleep(time.Second * 1)
		return "done 1"
	case "europe":
		//time.Sleep(time.Second * 2)
		return "done 2"
	case "america":
		//time.Sleep(time.Second * 5)
		return "done 3"
	}
	return "not found"
}
