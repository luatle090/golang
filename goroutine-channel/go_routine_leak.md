## Go Channel

ch := make(chan int)

Khai báo channel kiểu int

Channel có chức năng là send và receive (2 chiều). Send tức là gửi giá trị vào channel. Receive tức là nhận giá trị từ channel, từ góc độ của biến. 

**Lưu ý: Lưu ý góc độ của channel sẽ ngược, vì vậy ta xem channel như bưu điện. Biến x là người gửi, tới 1 biến a là người nhận**

ch <- x // send: x send value vào channel 

a := <-ch //receive: a nhận value từ channel

<-ch // receive and discard: nhận giá trị nhưng ko xử lý

### Khai báo channel 1 chiều

Khai báo này dành cho parameter hoặc signature return của hàm. 

chan<- int // send-only: giá trị x sẽ send vào channel

<-chan int // receive-only: x nhận giá trị từ channel

Về căn bản thì channel khi pass vào parameter thì là 2 chiều nhưng khi sử dụng ở trong hàm là 1 chiều như ví dụ bên dưới: channel natural là 2 chiều trước khi pass vào hàm counter nhưng trong hàm counter thì channel này chỉ sử dụng 1 chiều. Tương tự như vậy đối với hàm printer, trước khi pass là 2 chiều nhưng trong printer chỉ còn 1 chiều

Với signature return của hàm cũng vậy: Đôi lúc bạn khai báo muốn khai báo channel bên trong hàm thì lúc này channel là 2 chiều nhưng khi trả ra chỉ muốn người sử dụng sử dụng nó 1 chiều thì khai báo nó là 1 chiều.
Ví dụ: hàm Tick trong package time: func Tick(d Duration) <-chan Time


```Go
func counter (out chan<- int){
	x := 1 
	out <- x 	// send-only: biến x send giá trị vào channel 
	close(out)
}

func printer(in <-chan int){
	for {
		x, ok := <-in  	// receive-only: biến x nhận giá trị từ channel
		if !ok {
			break
		}
		fmt.Print(x)
	}
}

func main(){
	natural := make(chan int)
	go counter(natural)
	printer(natural)
}
```

-------

## Leak

```go
package main

import "fmt"

func counter(out chan<- int) {
	for i := 0; i < 100; i++ {
		out <- i
	}
	fmt.Println("closed counter")
	close(out)
}

func computing(in <-chan int, out chan<- int) {
	for v := range in {
		out <- v * v
	}
	fmt.Println("closed computing")
	close(out)
}

func printer(in <-chan int) {
	for v := range in {
		fmt.Printf("%v ", v)
	}
}

func main() {

	naturals := make(chan int)
	squares := make(chan int)

	go counter(naturals)
	go computing(naturals, squares)

	printer(naturals)
}
```
#### Ví dụ 1: rò rỉ do pass sai channel
giá trị in ra là: 
```cmd
0 2 3 4 5 6 7 8 9 10 ... 97 98 99 close counter
```
Để ý: giá trị 1 bị mất trong dãy in ra. Và dòng print thông báo ```closed computing``` ko xuất hiện.
**Lý do:** Giá trị 1 đã đc nhận trong computing nhưng khi pass vào biến out trong computing thì ko có giá trị nào nhận tiếp nên channel này bị chặn lại, đó cũng là lý do closed computing ko xuất hiện trong chương trình. => goroutine leak

#### Ví dụ 2: Đổi chỗ channel go 

```go
func main() {
	naturals := make(chan int)
	squares := make(chan int)

	go counter(naturals)
	go computing(squares, naturals)

	printer(naturals)
}
```

Hàm computing(squares, nauturals) đổi chỗ channel.
giá trị in ra là 0 1 2 3 4 5 6 7 8 9 10 ... 97 98 99 close counter
Chỉ close mỗi counter.
**Lý do:** do channel đầu vào là squares mà channel này hiện ko có giá trị nào được gửi vào nên goroutine này sẽ bị block cho đến khi có giá trị được gửi vào. Hàm printer thì tiêu thụ naturals nên khi tiêu thị hết thì goroutine counter có thể đóng được luồng.

---
# Leak goroutine (dịch từ blog uber)

Một thành phần quan trọng của ngữ nghĩa channel là blocking, trong đó hoạt động của channel tạm dừng thực thi goroutine cho đến khi đạt được điểm hẹn (nghĩa là tìm thấy đối tác liên lạc). 
Cụ thể hơn, đối với một channel không có bộ đệm, người gửi sẽ bị chặn cho đến khi có người nhận channel và ngược lại. Một goroutine có thể bị chặn mãi mãi khi cố gửi hoặc nhận trên một channel; tình huống mà một goroutine bị chặn không bao giờ được bỏ chặn được gọi là rò rỉ goroutine . Khi quá nhiều goroutine bị rò rỉ, hậu quả có thể rất nghiêm trọng. Các goroutine bị rò rỉ tiêu thụ các tài nguyên như bộ nhớ không được giải phóng hoặc thu hồi. 

> **Lưu ý**: Các channel đệm cũng có thể gây rò rỉ goroutine khi bộ đệm đầy. Các goroutine gửi khác sẽ bị block lại cho cho đến khi buffer còn vị trí trống. 

Các lỗi lập trình (ví dụ: luồng điều khiển phức tạp, trả về sớm, timeout), có thể dẫn đến sự không khớp trong giao tiếp giữa các goroutine, trong đó một hoặc nhiều goroutine có thể bị chặn nhưng không có goroutine nào khác đc tạo điều kiện cần thiết để bỏ chặn. Rò rỉ goroutine ngăn bộ thu gom rác thu hồi channel được liên kết, ngăn xếp goroutine và tất cả các đối tượng có thể truy cập của goroutine bị chặn vĩnh viễn. Đối với các dịch vụ hoạt động lâu dài, các rò rỉ nhỏ sẽ tích tụ theo thời gian sẽ làm trầm trọng thêm vấn đề. 

### Giải pháp thiết thực, gọn nhẹ để giải quyết rò rỉ Goroutine trong quá trình lập trình

Chúng tôi thực hiện một cách tiếp cận thực tế để phát hiện rò rỉ goroutine đối với các chương trình chạy lâu dài trong quá trình sản xuất đáp ứng tiêu chí đã nói ở trên. Tiền đề và quan sát quan trọng của chúng tôi như sau:

1. Nếu một chương trình có số lượng rò rỉ goroutine không nhỏ, thì cuối cùng nó sẽ được biểu thị thông qua số lượng lớn goroutine bị chặn trên một số hoạt động của channel
2. Chỉ một vài vị trí source code (liên quan đến hoạt động của channel) chiếm hầu hết các goroutine bị rò rỉ
3. Rò rỉ goroutine hiếm gặp, trong khi không lý tưởng, phát sinh chi phí thấp và có thể bị bỏ qua

Với quan quan điểm 1 được chứng minh bằng sự gia tăng đột biến về số lượng goroutine cho các chương trình bị rò rỉ. 

Với quan điểm #2 chỉ đơn giản nói rằng không phải tất cả các hoạt động của channel đều bị rò rỉ, nhưng vị trí nguồn của các nguyên nhân rò rỉ đáng kể phải được thực hiện thường xuyên để phát hiện rò rỉ. Vì bất kỳ goroutine bị rò rỉ nào vẫn tồn tại trong phần còn lại của vòng đời dịch vụ, nên việc liên tục gặp phải các tình huống rò rỉ cuối cùng sẽ góp phần tạo ra một lượng lớn goroutine bị chặn. Điều này đặc biệt đúng đối với các rò rỉ gây ra bởi các thao tác đồng thời có thể truy cập thông qua nhiều đường dẫn thực thi khác nhau và/hoặc trong các vòng lặp. 

Cuối cùng, điểm #3 là một cân nhắc thực tế. Nếu một thao tác gây rò rỉ hiếm khi gặp phải, thì tác động của nó đối với việc tích lũy bộ nhớ sẽ ít nghiêm trọng hơn. Dựa trên những quan sát thực tế này, chúng tôi đã thiết kế LeakProf, một chỉ báo rò rỉ đáng tin cậy với ít thông báo sai và chi phí thời gian chạy tối thiểu.

### Các mẫu chương trình gây ra rò rỉ

1. Hàm trả về sớm

Mẫu này xảy ra bất cứ khi nào một số goroutine dự kiến sẽ giao tiếp, nhưng một số mã if lại return sớm mà không tham gia vào channel giao tiếp, khiến channel send phải chờ đợi mãi mãi (bị block do ko có gorountine nào nhận). Điều này xảy ra khi các đối tác truyền thông không tính đến tất cả các đường đi thực thi có thể có.

```go {linenos=table,hl_lines=[14]}
func f(...){
	c := make(chan error)
	go func() {
		if err != nil {
			c <- err
			return
		}
		c <- nil
	}()
	if .. {
		..
		if .. {
			return ..
		} else if .. {
			return ..
		}
	}
	err := <- c
}
```

Hàm này tạo kênh c được sử dụng cho goroutine con sau đó, và được sử dụng để gửi error vào c. Hoạt động trên luồng chính bắt đầu từ đoạn if và có các return có thể được trả về từ luồng chính. Nếu main goroutine đúng các điều kiện trong if, thì goroutine con sẽ bị block mãi mãi.

Một giải pháp khả thi để ngăn chặn rò rỉ goroutine trong đó goroutine con gửi tin nhắn trên kênh và parent goroutine có thể muốn bỏ qua là tạo kênh có kích thước bộ đệm là 1. Điều này cho phép người sub goroutine ko bị chặn, bất kể hành vi của goroutine receiver.

### Rò rỉ thời gian chờ (Timeout)

Lỗi này có thể được coi là một trường hợp đặc biệt của mẫu hàm return quá sớm.
Rò rỉ này thường xuất hiện khi kết hợp việc sử dụng kênh không có bộ đệm với bộ hẹn giờ hoặc ngữ cảnh và từ khóa select. Bộ đếm thời gian hoặc context thường được sử dụng để đoản mạch quá trình thực thi goroutine và hủy bỏ sớm. Tuy nhiên, nếu goroutine con không được thiết kế để xem xét tình huống này, nó có thể dẫn đến rò rỉ.

```go
func TimeoutBug(inputCtx contet.Context, ...){
	ctx, cancel := context.WithTimeout(inputCtx, ...)
	done := make(chan any)
	go func(){
		...
		done <- data
	}()
	select {
		case data := <- done:
			return data
		case <- ctx.Done()
			return ...
	}
}
```

Tình huống bên trên xảy ra khi context đã quá timeout, hàm sẽ return nhưng sub goroutine vẫn còn hoạt động và gửi kết quả về done. Hậu quả là gây ra rò rỉ.

Channel done được sử dụng với goroutine con (dòng 4). Khi goroutine con gửi một tin nhắn (dòng 6), nó sẽ chặn cho đến khi một goroutine khác receive từ done. Trong khi đó, goroutine cha đợi ở câu lệnh select trên dòng 8, cho đến khi nó đồng bộ hóa với con (dòng 9) hoặc khi ctx hết thời gian chờ (dòng 11). Trong trường hợp hết thời gian chờ theo ngữ cảnh, goroutines cha sẽ trả về thông qua trường hợp trên dòng 11; kết quả là không có người nhận nào đứng đợi khi sub goroutine gửi. Do đó, sẽ bị rò rỉ, vì không có goroutine nào khác sẽ nhận được từ done.

Rò rỉ này có thể được ngăn chặn tương tự như giải pháp trên bằng cách tăng buffer lên 1.


### Rò rỉ nCast

Mẫu này xảy ra bất cứ khi nào hệ thống đồng thời giao tiếp giữa nhiều người gửi và một người nhận trên cùng một kênh. Hơn nữa, nếu người nhận chỉ thực hiện nhận một lần trên channel, thì tất cả goroutine gửi trừ đi một goroutine đã gửi thành công sẽ chặn mãi mãi trên kênh.

leaks = all_sub_go_sender - sub_go_send_success

```go
func Communication(...){
	dataChan := make(chan any)
	for _, item := range items {
		go func(c chan any, ...){
			c <- result
		}(dataChan, ...)
	}
	result := <- dataChan
}

```

Kênh dataChan (dòng 2) được truyền dưới dạng đối số cho các goroutine được tạo trong for-loop trên dòng 4. Mỗi goroutine con cố gắng gửi kết quả đến dataChan, nhưng goroutine cha chỉ nhận từ dataChan một lần rồi thoát, tại thời điểm đó, nó sẽ mất tham chiếu đến dataChan. Vì dataChan không có bộ đệm, bất kỳ sub nào không đồng bộ hóa với parent sẽ bị chặn mãi mãi.

Giải pháp là tăng buffer của dataChan lên bằng độ dài của items. Điều này bảo toàn đc yêu cầu chỉ có kết quả đầu tiên được gửi tới dataChan sẽ được nhận bởi goroutine cha, đồng thời cho phép các goroutine con còn lại bỏ chặn và kết thúc.

Một dạng tổng quát hơn của vấn đề này xảy ra khi có N người gửi và M người nhận, N > M và mỗi người nhận chỉ thực hiện một thao tác nhận.

### Lạm dụng for-range

Kiểu rò rỉ này có thể xảy ra khi sử dụng cấu trúc for-range với các kênh. Để hiểu được rò rỉ này, bạn cần phải làm quen với thao tác đóng và cách hoạt động của for-range với các kênh. Rò rỉ xảy ra bất cứ khi nào một kênh được lặp lại, nhưng kênh đó không bao giờ bị đóng. Điều này khiến vòng lặp for lặp qua kênh bị chặn vô thời hạn vì vòng lặp không kết thúc trừ khi kênh bị đóng; các khối vòng lặp for sau khi nhận được tất cả các mục từ kênh.

```go
func ChannelIteration(){
	wg := &sync.WaitGroup{}
	queueJobs := make(chan any, 1)
	for ... := range jks {
		wg.Add(1)
		go func(){
			queueJobs <- data
		}()
	}

	// consume queue
	go func() {
		for data := range queueJobs {
			jobs = append(job, data)
			wg.Done()
		}
	}()
	wg.Wait()
}

```
Để cho ngắn gọn, chúng ta sẽ mượn kiến trúc producer-consumer được giới thiệu trong Tranh chấp truyền thông. Channel queueJobs được cấp phát ở dòng 3. Các producer là các goroutines được tạo ra trong vòng lặp for (dòng 3) với từ khóa go (dòng 5), trong đó mỗi người gửi gửi một thông báo (dòng 5). Consumer (dòng 10) nhận các message bằng cách dùng for-range qua queueJobs, vòng lặp ở dòng 11 sẽ thực hiện nhận message. Mong đợi ở đây là một khi producer không gửi message nữa, thì consumer sẽ thoát khỏi vòng lặp và kết thúc. Tuy nhiên, trong trường hợp không có thao tác đóng trên kênh, for-range sẽ bị chặn vì không có thêm message nào được gửi, dẫn đến rò rỉ.

**Giải pháp:**
Vì parent của producer-consumer đợi cho đến khi tất cả thông điệp được gửi (thông qua cấu trúc WaitGroup), nên giải pháp là thêm ```close(queueJobs)``` sau wg.Wait() (dòng 16) hoặc dưới dạng câu lệnh ```defer close(queueJobs)``` ngay sau khai báo queueJobs. Sau khi tất cả các tin nhắn được gửi, parent goroutine sẽ đóng queueJobs, báo hiệu cho consumer ngừng lặp queueJobs và do đó kết thúc và được thu gom rác.

Cách 1: Thêm ```close(queueJobs)``` sau wg.Wait()

```go
func ChannelIteration(...){
	wg := &sync.WaitGroup{}
	queueJobs := make(chan any, 1)
	for ... := range jks {
		wg.Add(1)
		go func(){
			queueJobs <- data
		}()
	}

	// consume queue
	go func() {
		for data := range queueJobs {
			jobs = append(job, data)
			wg.Done()
		}
	}()
	wg.Wait()
	close(queueJobs)
}
```

Cách 2: Sử dụng ```defer close(queueJobs)```

```go
func ChannelIteration(){
	wg := &sync.WaitGroup{}
	queueJobs := make(chan any, 1)

	defer close(queueJobs)
	
	for ... := range jks {
		wg.Add(1)
		go func(){
			queueJobs <- data
		}()
	}

	// consume queue
	go func() {
		for data := range queueJobs {
			jobs = append(job, data)
			wg.Done()
		}
	}()
	wg.Wait()
}
```

Tham khảo

https://www.uber.com/en-VN/blog/leakprof-featherlight-in-production-goroutine-leak-detection/?_gl=1*1p11hz5*_ga*MTM4NjkxMzE5My4xNzEwMjExMDI0*_ga_XTGQLY6KPT*MTcxMDIxMTAyNC4xLjEuMTcxMDIxMTAyNC4wLjAuMA..&_ga=2.182629990.1553125118.1710211025-1386913193.1710211024


# Sử dụng goleak để phát hiện rò rỉ

https://github.com/uber-go/goleak

```go
func leak() error{
	go func(){
		time.Sleep(time.Minute)
	}()
	return nil
}


// function test

func TestLeakFunction(t *testing.T){
	defer goleak.VerifyNone(t)
	if err := leak(); err != nil {
		t.Fatal("error")
	}
}
```

Thông báo cung cấp thông tin hữu ích trong thông báo lỗi:

- Stack trace của goroutine bị rò rỉ, cùng với trạng thái của goroutine. Thông tin này có thể giúp gỡ lỗi nhanh chóng và hiểu goroutine nào đang bị rò rỉ.
- ID goroutine, hữu ích khi trực quan hóa quá trình thực thi với trình theo dõi. Dưới đây là một ví dụ về các dấu vết được tạo bằng các bài kiểm tra

```console
go test ./path/to/file_test.go -trace trace.out
```

## Internal
Yêu cầu duy nhất để kích hoạt phát hiện rò rỉ là gọi thư viện vào cuối bài kiểm tra để phát hiện bất kỳ sub goroutine nào bị rò rỉ. Trên thực tế, nó kiểm tra bất kỳ sub goroutine bổ sung nào thay vì chỉ sub goroutine bị rò rỉ

1 số hạn chế:
- Dương tính giả