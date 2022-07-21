// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 254.
//!+

// Chat is a server that lets clients chat with each other.
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

//!+broadcaster
type client2 chan<- string // an outgoing message channel
type client struct {
	msg     chan<- string
	cliName string
}

var (
	timeout  = 10 * time.Second
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string) // all incoming client messages
)

func broadcaster() {
	clients := make(map[client]bool) // all connected clients
	for {
		select {
		case msg := <-messages:
			sendMessage(clients, msg)

		case cli := <-entering:
			sendMessage(clients, cli.cliName+" has arrived")
			clients[cli] = true

		case cli := <-leaving:
			delete(clients, cli)
			sendMessage(clients, cli.cliName+" has left")
			close(cli.msg)
		}
	}
}

//!-broadcaster

// Broadcast incoming message to all
// clients' outgoing message channels.
func sendMessage(clients map[client]bool, msg string) {
	for cli := range clients {
		cli.msg <- msg
	}
}

func idle(conn net.Conn, timer *time.Timer, reset <-chan struct{}, exit <-chan struct{}) {
	for {
		select {
		case <-timer.C:
			fmt.Println("idle 10 seconds")
			conn.Close()
			return
		case <-reset:
			timer.Reset(timeout)
		case <-exit:
			conn.Close()
			return
		}
	}
}

//!+handleConn
func handleConn(conn net.Conn) {
	ch := make(chan string) // outgoing client messages
	reset := make(chan struct{})
	exit := make(chan struct{})

	// open output
	go clientWriter(conn, ch)

	who := conn.RemoteAddr().String()
	ch <- "You are " + who
	cli := client{ch, who}
	entering <- cli

	// computation idle
	timer := time.NewTimer(timeout)
	go idle(conn, timer, reset, exit)
	input := bufio.NewScanner(conn)
	for input.Scan() {
		messages <- who + ": " + input.Text()
		reset <- struct{}{}
	}

	// nil occur when user types Control-C
	fmt.Println(input.Err())
	if input.Err() == nil {
		fmt.Println("sssadas")
		exit <- struct{}{}
	}

	fmt.Println(who + ": has left")
	leaving <- cli
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg) // NOTE: ignoring network errors
	}
}

//!-handleConn

//!+main
func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	go broadcaster()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

//!-main
