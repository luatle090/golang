package main

import (
	"fmt"
	"sync"
)

func main() {
	Variant3()
}

func CommonConcurrencyPattern() {
	sc := make(chan []string, 3)

	sc <- []string{"a", "b"}
	sc <- []string{"a", "bss"}
	sc <- []string{"a", "baa"}
	close(sc)

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		//fmt.Println(<-sc)
		for sliceString := range sc {
			for _, word := range sliceString {
				fmt.Printf("%s ", word)
			}
			fmt.Println()
		}
		defer wg.Done()
	}()
	wg.Wait()
}

// doesn't using WaitGroup
func Variant2() {
	sc := make(chan []string, 3)

	sc <- []string{"a", "b"}
	sc <- []string{"a", "bss"}
	sc <- []string{"a", "baa"}
	close(sc)

	wait := make(chan struct{})
	go func() {
		//fmt.Println(<-sc)
		for sliceString := range sc {
			for _, word := range sliceString {
				fmt.Printf("%s ", word)
			}
			fmt.Println()
		}
		wait <- struct{}{}
	}()
	<-wait
}

// doesn't using the built-in close function
// it's accessing to buffered
func Variant3() {
	sc := make(chan []string, 3)

	sc <- []string{"a", "b"}
	sc <- []string{"a", "bss"}
	sc <- []string{"a", "baa"}

	wait := make(chan struct{})
	go func() {
		var i int
		for i = 0; i < 3; i++ {
			sliceString := <-sc
			for _, word := range sliceString {
				fmt.Printf("%s ", word)
			}
			fmt.Println()
		}
		wait <- struct{}{}
	}()
	<-wait
}
