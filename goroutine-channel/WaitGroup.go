package main

import (
	"fmt"
	"sync"
)

// dùng con trỏ vì nếu dùng value sẽ chỉ copy wg
// và ko thay đổi giá trị của counter trong wg
// xem lại gọi hàm khi giá trị là value trong sách
func runner1(wg *sync.WaitGroup) {
	defer wg.Done() // This decreases counter by 1
	fmt.Print("\nI am first runner")

}

// dùng con trỏ vì nếu dùng value sẽ chỉ copy wg
// và ko thay đổi giá trị của counter trong wg
// xem lại gọi hàm khi giá trị là value trong sách
func runner2(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Print("\nI am second runner")
}

func execute() {
	wg := new(sync.WaitGroup)
	wg.Add(2)

	// We are increasing the counter by 2
	// because we have 2 goroutines
	go runner1(wg)
	go runner2(wg)

	// This Blocks the execution
	// until its counter become 0
	wg.Wait()
}

func main() {
	// Launching both the runners
	execute()
}
