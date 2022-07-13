package main

func main() {
	deadlockBufferedSend()
}

func deadlockReceive() {

	res := make(chan int)
	<-res
}

func deadlockSend() {

	res := make(chan int)
	res <- 12
}

func deadlockBufferedReceive() {

	res := make(chan int, 1)
	<-res
}

// ko có receive khi buffered đã vượt quá capacity(sức chứa)
// hành động put thêm vào buffered sẽ làm cho send channel
// block cho đến khi buffered empty hoặc có receive
func deadlockBufferedSend() {

	res := make(chan int, 1)
	res <- 12
	res <- 12
}

// deadlock send chạy trước receive => block forever => deadlock
// solution: cho send vào 1 goroutine hoặc cho go receive chạy trước
func deadlockSend2() {

	res := make(chan int)
	res <- 12
	go func() {
		//time.Sleep(time.Second * 1)

		<-res
	}()
}

// deadlock receive chạy trước send => block forever => deadlock
// solution: cho receive vào 1 goroutine hoặc cho go send chạy trước receive
func deadlockReceive2() {

	res := make(chan int)
	<-res
	go func() {
		//time.Sleep(time.Second * 1)

		res <- 12
	}()
}
