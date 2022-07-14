// Reverb2 is a TCP server that simulates an echo.
// addition the disconnect feature
// will disconnects any client that shouts nothing within 10 seconds
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

func scanInput(input *bufio.Scanner, lines chan<- string, exit chan<- struct{}) {

	for input.Scan() {
		lines <- input.Text()
	}
	fmt.Println("exit goroutine")
	exit <- struct{}{}
}

//!+
func handleConn(c *net.TCPConn) {
	lines, exit := make(chan string), make(chan struct{})
	input := bufio.NewScanner(c)
	var wg *sync.WaitGroup = new(sync.WaitGroup) // create pointer WaitGroup

	defer func() {
		wg.Wait()
		c.Close()
	}()

	go scanInput(input, lines, exit)

	for {
		select {
		case <-time.After(10 * time.Second):
			fmt.Println(c, "close connections")
			return
		case <-lines:
			wg.Add(1)
			go echo(c, input.Text(), 1*time.Second, wg)
		case <-exit:
			return
		}
	}
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
