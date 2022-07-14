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
var tokens = make(chan struct{}, 20)

func crawl(url string) []string {
	fmt.Println(url)
	tokens <- struct{}{} // acquire a token
	list, err := links.Extract(url)
	<-tokens // release the token

	if err != nil {
		log.Print(err)
	}
	return list
}

//!-sema

//!+
func main() {
	worklist := make(chan []string)
	var n int = -1 // number of pending sends to worklist
	var depth int = 2

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

				go func(link string, n int) {
					defer wg.Done()
					wl := crawl(link)

					if n < depth-1 {
						worklist <- wl
					}

				}(link, n)
			}
		}

	}
	//close(worklist)
	wg.Wait()

}

//!-
