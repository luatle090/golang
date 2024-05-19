## Context

Sử dụng để hủy thao tác khác mà này thao tác này khi phản hồi lại ko còn sử dụng, giúp tiết kiệm tài nguyên.

### Ngữ cảnh

Ví dụ: 1 user gọi tới web của bạn để truy vấn giỏ hàng, nhưng vì lý do gì đó người user này tắt trình duyệt thì lúc này kết quả truy vấn trả về ko còn được sử dụng. Do đó nếu tiếp tục truy vấn thì sẽ gây lãng phí tài nguyên. Context giúp giải quyết việc này.

## Quy tắc trong standard library Go

Đặt tham số context là tham số đầu tiên của hàm

```go
func doSomething(ctx context.Context, x int32) {
	fmt.Println("Doing something!")
}

```

## Context TODO
Tạo ra empty context. Sử dụng như 1 placeholder nếu như bạn ko biết chắc về context sử dụng như nào

**Sử dụng tạm trong quá trình phát triển. Production code ko nên include context.TODO()**

```go
context.TODO()
```

```go
package main

import (
	"context"
	"fmt"
)

func doSomething(ctx context.Context) {
	fmt.Println("Doing something!")
}

func main() {
	ctx := context.TODO()
	doSomething(ctx)
}
```

Hàm doSomething nhận vào context. Tuy nhiên hàm này chưa sử dụng context để làm gì !!! 

Kết quả khi chạy

```cmd
Output
Doing something!
```

## Context background

Tương tự như TODO(). Tuy nhiên lúc này bạn đã biết cách sử dụng context vào mục đích nào 

```
ctx := context.Background()
```

## Http Server

Sử dụng context trong http request, có 2 cách:

1. Nhận context từ request

2. Wrap context vào request thông qua request.WithContext(ctx). Lúc này ta đã tạo ra new request mới có context.

**Khi sử dụng chỉ cần 1 trong 2 cách tùy theo mục đích sử dụng**

### Mẫu khai báo context của http

Hàm này mục đích chỉ khai báo cho 2 cách trên. Sử dụng 2 cách 1 là nhận context từ request, 2 là wrap vào request và nhận request mới với context.

```go
func Middleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()			// nhận request từ context
		r = r.WithContext(ctx)    	// wrap context vào request
		handler.ServeHTTP(w, r)
	})
}
```

### Cách 1: Nhận context từ request

```go
func Handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	// lấy query string
	gioHang := r.URL.Query().Get("type")
	if gioHang == "" {
		w.WriteHeader(http.StatusFound)
		w.Write([]byte("query string ko tồn tại"))
	}
	result, _ := doSomething(ctx, gioHang)
	// convert int to string and then to byte array
	w.Write([]byte(strconv.Itoa(result)))

}

func doSomething(ctx context.Context, data string) (int, error) {
	// xử lý nghiệp vụ với data
	return 100, nil
}

```

## Cách 2: Tạo request với context

```go
type ServiceCaller struct {
	client *http.Client
}

func (sc ServiceCaller) callAnotherService(ctx context.Context, data string)
(string, error) {
	req, err := http.NewRequest(http.MethodGet,
	"http://example.com?data="+data, nil)
	
	if err != nil {
		return "", err
	}
	
	req = req.WithContext(ctx)
	resp, err := sc.client.Do(req)
	
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("Unexpected status code %d",
		resp.StatusCode)
	}
		// do the rest of the stuff to process the response
	id, err := processResponse(resp.Body)
	return id, err
}
```


## Pass data vào trong context

Thêm dữ liệu vào context và chuyển context từ hàm này sang hàm khác, mỗi layer có thể thêm thông tin bổ sung. Ví dụ: hàm đầu tiên có thể thêm username vào context. Hàm tiếp theo có thể thêm file path vào context mà người dùng này muốn truy cập. Cuối cùng, hàm thứ ba sau đó có thể đọc tập tin từ đĩa hệ thống và ghi lại xem nó có được load thành công hay không cũng như người dùng nào đã cố load nó.

### Sử dụng

Sử dụng hàm ```context.WithValue()``` để pass value. Hàm này nhận 3 tham số. Gồm context, key, value. Tham số key, value nhận vào kiểu dữ liệu là any

Để get value sử dụng context.Value(). Hàm nhận 1 tham số là key



```go
func doSomething(ctx context.Context) {
	// nhận  dữ liệu qua Value()
	fmt.Printf("doSomething: myKey's value is %s\n", ctx.Value("key"))
}

func main() {
	ctx := context.Background()

	// pass context, key, value vào hàm WithValue()
	ctx = context.WithValue(ctx, "key", "myValue")

	doSomething(ctx)
}
```

output

```cmd
doSomething: key's value is myValue
```

### Context là immutable Lưu ý khi pass

Trong một chương trình lớn hơn chạy trên máy chủ, giá trị này có thể là thời điểm mà chương trình bắt đầu chạy.

Khi sử dụng context, **điều quan trọng cần biết là các giá trị được lưu trữ trong một ngữ cảnh cụ thể context.Context là immutable(không thay đổi)**. Khi bạn gọi context.WithValue, bạn đã chuyển vào parent context và bạn cũng nhận lại được bản copy của parent. Vì hàm context.WithValue này không sửa đổi ngữ cảnh mà bạn cung cấp. Thay vào đó, nó bao bọc parent context của bạn bên trong một context khác với giá trị mới.

Hàm bên dưới đầu tiên tạo ra function look up key trong context được pass vào. Sau đó tạo key, tìm key như việc sử dụng bình thường. Tiếp tục khi bạn muốn thay đổi value của key "lnguage" thành là "C" thì WithValue sẽ bọc ctx cũ vào ctx mới và tạo ra context mới ctx2. Khi query ctx2 thì giá trị là C, truy vấn lại ctx cũ thì giá trị vẫn là "Go"

```go
func main() {

	//
	type favContextKey string

	// anonymous function
	f := func(ctx context.Context, k favContextKey) {
		if v := ctx.Value(k); v != nil {
			fmt.Println("found value:", v)
			return
		}
		fmt.Println("key not found:", k)
	}

	k := favContextKey("language")
	ctx := context.WithValue(context.Background(), k, "Go")

	// pass vào key
	f(ctx, k)
	f(ctx, favContextKey("color"))

	// 
	ctx2 := context.WithValue(ctx, k, "C")
	f(ctx2, k)
	f(ctx, k)

}
```

output
```cmd
found value: Go
key not found: color
found value: C
found value: Go
```

**Lưu ý:** Docs go đưa ra lời khuyên là key phải là giá trị có thể so sánh được, ko nên là kiểu string hoặc bất kỳ kiểu có sẵn nào để tránh conflict giữa các package sử dụng context. Vì vậy user nên định nghĩa ra kiểu riêng như dòng đầu tiên trong main ở hàm trên. Để tránh việc phân bổ khi gán cho một interface{}, context keys thường là kiểu struct cụ thể struct{}. Alternatively, exported context key variables' static type should be a pointer or interface.

### Cạm bẫy

*Ngữ cảnh có thể là một công cụ mạnh mẽ với tất cả các giá trị mà chúng có thể chứa, nhưng cần phải đạt được sự cân bằng giữa dữ liệu được lưu trữ trong ngữ cảnh và dữ liệu được truyền đến hàm dưới dạng tham số. Việc đặt tất cả dữ liệu của bạn vào một ngữ cảnh và sử dụng dữ liệu đó trong các hàm thay vì tham số có vẻ hấp dẫn, nhưng điều đó có thể dẫn đến mã khó đọc và khó bảo trì. Một nguyên tắc chung là mọi dữ liệu cần thiết để hàm chạy phải được chuyển dưới dạng tham số. 
Ví dụ: việc giữ các giá trị như tên người dùng trong giá trị ngữ cảnh để sử dụng khi ghi thông tin sau này có thể hữu ích. Tuy nhiên, nếu tên người dùng được sử dụng để xác định xem một hàm có hiển thị một số thông tin cụ thể hay không, bạn sẽ muốn đưa nó vào làm tham số hàm ngay cả khi nó đã có sẵn trong ngữ cảnh. Bằng cách này, khi bạn hoặc người khác xem xét chức năng này trong tương lai, bạn sẽ dễ dàng biết được dữ liệu nào đang thực sự được sử dụng.*

## Hủy context

## Xác định context is done

