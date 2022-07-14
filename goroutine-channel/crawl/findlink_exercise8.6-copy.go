package main

import (
	"fmt"
	"log"
	"os"
	"sync"

	"gopl.io/ch5/links"
)

//!+sema
// tokens is a counting semaphore used to
// enforce a limit of 20 concurrent requests.
var tokens = make(chan struct{}, 3)

func crawl(url string) []string {
	tokens <- struct{}{} // acquire a token
	list, err := links.Extract(url)
	<-tokens // release the token

	if err != nil {
		log.Print(err)
	}
	return list
}

func print(url string, n int) {
	fmt.Println(n+1, url)
}

//!-sema

//!+
func main() {
	worklist := make(chan []string)
	var n int = 0 // number of pending sends to worklist

	depth := 1
	var wg sync.WaitGroup

	// Start with the command-line arguments.
	go func() { worklist <- os.Args[1:] }()

	// Crawl the web concurrently.
	seen := make(map[string]bool)
	for ; n < depth; n++ {
		list := <-worklist
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				wg.Add(1)
				if n < depth {
					n--
				}
				go func(link string, n int) {
					print(link, n)

					if n < depth {
						worklist <- crawl(link)
					}
					defer wg.Done()
				}(link, n)
			}
		}

	}

	wg.Wait()
}

//!-
