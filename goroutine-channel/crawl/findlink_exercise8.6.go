// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 243.

// Crawl3 crawls web links starting with the command-line arguments.
//
// This version uses bounded parallelism.
// For simplicity, it does not address the termination problem.
//

package main

import (
	"fmt"
	"log"
	"os"
	"sync"

	"gopl.io/ch5/links"
)

func crawl(url string) []string {
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}

//!+
func main() {
	worklist := make(chan []string)  // lists of URLs, may have duplicates
	unseenLinks := make(chan string) // de-duplicated URLs

	var wg sync.WaitGroup
	var n, depth int
	depth = 1
	// Add command-line arguments to worklist.
	go func() { worklist <- os.Args[1:] }()

	// Create 20 crawler goroutines to fetch each unseen link.
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func() {
			for link := range unseenLinks {
				foundLinks := crawl(link)
				// fmt.Println(j)
				// worklist <- foundLinks
				// j++

				go func() { worklist <- foundLinks }()

			}
			defer wg.Done()
		}()
	}

	// The main goroutine de-duplicates worklist items
	// and sends the unseen ones to the crawlers.
	seen := make(map[string]bool)

	go func() {
		wg.Wait()
		fmt.Println("lll")
		close(unseenLinks)

	}()

	for list := range worklist {

		//fmt.Println("list ", list)
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				unseenLinks <- link
			}
		}
		fmt.Println()
		if n == depth {
			fmt.Println("11")
			close(worklist)
		}
		n++
	}

}


