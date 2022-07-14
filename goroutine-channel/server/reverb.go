// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 224.

// Reverb2 is a TCP server that simulates an echo.
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"sync"
	"time"
)

// Dùng con trỏ vì nếu dùng value sẽ chỉ copy wg
// và ko thay đổi giá trị của counter trong wg
// xem lại gọi hàm khi giá trị là value trong sách
func echo(c net.Conn, shout string, delay time.Duration, wg *sync.WaitGroup) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
	defer wg.Done()
}

//!+
func handleConn(c *net.TCPConn) {
	input := bufio.NewScanner(c)
	var wg *sync.WaitGroup = new(sync.WaitGroup) // create pointer WaitGroup
	for input.Scan() {
		wg.Add(1)
		go echo(c, input.Text(), 1*time.Second, wg)
	}

	wg.Wait()
	c.CloseWrite()

}

//!-

func main() {
	tcpAddr, errTCPAddr := net.ResolveTCPAddr("tcp", "localhost:8000")
	if errTCPAddr != nil {
		log.Fatal(errTCPAddr)
	}
	l, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := l.AcceptTCP()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		go handleConn(conn)
	}

}
