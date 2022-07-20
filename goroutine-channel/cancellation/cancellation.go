package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

var done = make(chan struct{})
var max, min int = 10, 1

func main() {
	res := make(chan string)
	var n sync.WaitGroup
	rand.Seed(time.Now().UnixNano())
	req := [3]string{"asia", "europe", "america"}

	for i := 0; i < 100; i++ {
		n.Add(1)
		go fetch(req[rand.Intn(3)], res, &n)
	}

	go func() {
		n.Wait()
		fmt.Println("close goroutines")
		close(res)
	}()

	var s string

	s = <-res
	close(done)
	fmt.Println(s)

	for range res {

	}

}

func fetch(hostname string, res chan<- string, n *sync.WaitGroup) {
	res <- mirroredQueue(hostname)
	defer n.Done()
}

var sema = make(chan struct{}, 20)

func mirroredQueue(hostname string) string {
	select {
	case sema <- struct{}{}: // acquire token
	case <-done:
		return "" // cancelled
	}
	defer func() { <-sema }()

	var reqTime = rand.Intn(max) + min
	time.Sleep(time.Second * time.Duration(reqTime))
	var result = strconv.Itoa(reqTime)
	switch hostname {
	case "asia":
		return result + " done 1"
	case "europe":
		return result + " done 2"
	case "america":
		return result + " done 3"
	}
	return "not found"
}
