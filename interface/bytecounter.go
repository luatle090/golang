package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type ByteCounter int

func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p)) // convert int to ByteCounter
	return len(p), nil
}

func main() {
	var c ByteCounter
	c.Write([]byte("hello"))
	fmt.Println(c)

	c = 0
	var name = "Dolly"
	fmt.Fprintf(&c, "hello, %s", name)
	fmt.Println(c) // "12", = len("hello, Dolly")

	var w WordCounter
	input := "Now is the winter of our discontent,\nMade glorious summer by this sun of York.\n"
	w.Write([]byte(input))
	fmt.Println("word counter", w)

	var line LineCounter
	line.Write([]byte(input))
	fmt.Println("Line counter", line)
}

type WordCounter int

func (l *WordCounter) Write(p []byte) (int, error) {
	scanner := bufio.NewScanner(strings.NewReader(string(p)))
	scanner.Split(bufio.ScanWords)
	count := 0
	for scanner.Scan() {
		count++
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}
	*l += WordCounter(count)
	return count, nil
}

type LineCounter int

func (l *LineCounter) Write(p []byte) (int, error) {
	scanner := bufio.NewScanner(strings.NewReader(string(p)))
	scanner.Split(bufio.ScanLines)
	count := 0
	for scanner.Scan() {
		count++
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}
	*l += LineCounter(count)
	return count, nil
}

/*
package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

type byteCounter int
type wordCounter int
type lineCounter int

func main() {
    var c byteCounter
    c.Write([]byte("Hello This is a line"))
    fmt.Println("Byte Counter ", c)

    var w wordCounter
    w.Write([]byte("Hello This is a line"))
    fmt.Println("Word Counter ", w)

    var l lineCounter
    l.Write([]byte("Hello \nThis \n is \na line\n.\n.\n"))
    fmt.Println("Length ", l)

}

func (c *byteCounter) Write(p []byte) (int, error) {
    *c += byteCounter(len(p))
    return len(p), nil
}

func (w *wordCounter) Write(p []byte) (int, error) {
    count := retCount(p, bufio.ScanWords)
    *w += wordCounter(count)
    return count, nil
}

func (l *lineCounter) Write(p []byte) (int, error) {
    count := retCount(p, bufio.ScanLines)
    *l += lineCounter(count)
    return count, nil
}

func retCount(p []byte, fn bufio.SplitFunc) (count int) {
    s := string(p)
    scanner := bufio.NewScanner(strings.NewReader(s))
    scanner.Split(fn)
    count = 0
    for scanner.Scan() {
        count++
    }
    if err := scanner.Err(); err != nil {
        fmt.Fprintln(os.Stderr, "reading input:", err)
    }
    return
}

*/
