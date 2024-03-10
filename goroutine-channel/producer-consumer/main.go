package main

import (
	"fmt"
	"sync"
)

// struct cho producer consumer
type PubCon struct {
	out chan any
}

// Sử dụng struct producer consumer
type ServiceX struct {
	*PubCon
	wg *sync.WaitGroup
}

func CreatePubCon() *PubCon {
	return &PubCon{
		out: make(chan any),
	}
}

func (pc PubCon) Producer(data any) {
	pc.out <- data
}

func (pc PubCon) Close() {
	close(pc.out)
}

func (pc PubCon) Consumer() <-chan any {
	return pc.out
}

// Hàm phát ra counter rồi pass vào channel thông qua producer
func (sc ServiceX) Counter() {
	ten := make([]int, 10)
	defer sc.wg.Done()
	counter := 1
	for i := range ten {
		sc.Producer(counter)
		counter++
		ten[i] = counter
	}
}

// Hàm nhận data từ channel via Consumer
func (sc ServiceX) Printer() {
	r := sc.Consumer()
	for x := range r {
		i := x.(int)
		fmt.Println(i)
	}
}

// producer - consumer đơn giản
func main(){
	pc := CreatePubCon()
	// assert.NotEmpty(pc, "pc is nil")
	if pc == nil {
		fmt.Println("struct Producer Consumer is nil")
		return
	}
	sx := ServiceX{
		PubCon: pc,

		// Cấp phát vùng nhớ cho WaitGroup trong struct!!!
		wg:     new(sync.WaitGroup),
	}

	sx.wg.Add(1)
	go sx.Counter()

	// close channel when counter goroutine was done
	go func() {
		sx.wg.Wait()
		sx.Close()
	}()
	sx.Printer()
}