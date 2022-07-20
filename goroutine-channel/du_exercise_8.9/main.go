// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 250.

// The du3 command computes the disk usage of the files in a directory.
package main

// The du3 variant traverses all directories in parallel.
// It uses a concurrency-limiting counting semaphore
// to avoid opening too many files at once.

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"time"
)

var vFlag = flag.Bool("v", false, "show verbose progress messages")

type FileSize struct {
	root   string
	size   int64
	nFiles int
}

//!+
func main() {
	// ...determine roots...

	//!-
	flag.Parse()

	// Determine the initial directories.
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}

	//!+
	// Traverse each root of the file tree in parallel.
	var mapSize = make(map[string]FileSize)
	fileSizes := make(chan FileSize)
	var n sync.WaitGroup
	for _, root := range roots {
		n.Add(1)
		mapSize[root] = FileSize{root: root}
		go walkDir(root, root, &n, fileSizes)
	}
	go func() {
		n.Wait()
		close(fileSizes)
	}()
	//!-

	// Print the results periodically.
	var tick <-chan time.Time

	if *vFlag {
		tick = time.Tick(100 * time.Millisecond)
	}

loop:
	for {
		select {
		case fileSize, ok := <-fileSizes:
			if !ok {
				break loop // fileSizes was closed
			}
			file := mapSize[fileSize.root]
			file.size += fileSize.size
			file.nFiles++
			mapSize[file.root] = file
		case <-tick:
			printDiskUsage(mapSize)
		}
	}

	printDiskUsage(mapSize) // final totals
}

func printDiskUsage(mapSize map[string]FileSize) {
	for root, file := range mapSize {
		fmt.Printf("%10d files  %.1f GB  under %s\n", file.nFiles, float64(file.size)/1e9, root)
	}
}

// walkDir recursively walks the file tree rooted at dir
// and sends the size of each found file on fileSizes.
//!+walkDir
func walkDir(dir string, root string, n *sync.WaitGroup, fileSizes chan<- FileSize) {
	defer n.Done()
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			n.Add(1)
			subdir := filepath.Join(dir, entry.Name())
			go walkDir(subdir, root, n, fileSizes)
		} else {
			fileSizes <- FileSize{root: root, size: entry.Size()}
		}
	}
}

//!-walkDir

//!+sema
// sema is a counting semaphore for limiting concurrency in dirents.
var sema = make(chan struct{}, 20)

// dirents returns the entries of directory dir.
func dirents(dir string) []os.FileInfo {
	sema <- struct{}{}        // acquire token
	defer func() { <-sema }() // release token
	// ...
	//!-sema

	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du: %v\n", err)
		return nil
	}
	return entries
}
