## Interface căn bản

https://www.golangprograms.com/go-language/interface.html

Định nghĩa tập các phương thức. Giá trị của _kiểu interface_ có thể giữ bất kỳ giá trị. 

Để phương thức thỏa mãn interface thì bắt buộc kiểu đó phải implement tất cả các phương thức đc định nghĩa trong interface.

Interface ```Employee``` Định nghĩa hàm in ```PrintSalary```. Để implement interface thì kiểu ```Emp``` ở hàm main đã thực hiện việc gán biến, như vậy là đã implement interface

**Lưu ý**: Để Implement đc Interface thì ko đc khai báo interface là con trỏ ```var e1 *Employee```

```go
func main(){
	var e1 Employee
	e1 = Emp(1)
	fmt.Println("Employee Salary:", e1.PrintSalary(25000, 5))
}
```

```go
package main

import "fmt"

type Employee interface {
	PrintSalary(basic int, tax int) int
}

type Emp int

// PrintSalary method to calculate employee salary
func (e Emp) PrintSalary(basic int, tax int) int {
	var salary = (basic * tax) / 100
	return basic - salary
}

func main() {
	var e1 Employee
	e1 = Emp(1)
	e1.PrintName("John Doe")
	fmt.Println("Employee Salary:", e1.PrintSalary(25000, 5))
}
```

## Interface là implicitly (ngầm định)

Trong Go ko có từ khóa tường minh nào để chỉ ý định là "implements". Việc implement chỉ cần type (từ khóa) có hàm có cấu trúc giống như cấu trúc của interface thì sẽ implement.

Vì interface ko có từ khóa "implements" nên mọi phương thức implement  các phương thức được định nghĩa của interface gọi là thỏa mãn.

Để phương thức thỏa mãn interface thì bắt buộc kiểu đó phải implement tất cả các phương thức đc định nghĩa trong interface.

Định nghĩa I interface có hàm M(). Để Kiểu T implement hàm đó thì cần viết hàm giống như M(). Sau đó gán biến sang interface như ở hàm main
```go
func (t T) M() {
	fmt.Println(t.S)
}
```

```golang
package main

import "fmt"

type I interface {
	M()
}

type T struct {
	S string
}

// This method means type T implements the interface I,
// but we don't need to explicitly declare that it does so.
func (t T) M() {
	fmt.Println(t.S)
}

func main() {
	var i I = T{"hello"}
	i.M()
}
```

## Interface chấp nhận con trỏ

**Nếu method là type value thì interface nhận type value. Nếu method là receiver pointer thì interface phải nhận địa chỉ.**

Print() method của type Book bên dưới là receiver pointer. Do đó, để thỏa mãn thì interface phải accept receiver pointer.

Đối với Print() của type Magazine sẽ là type value. Nên việc gán địa chỉ hay ko địa chỉ ko quan trọng.

_Thử xóa bỏ địa chỉ của biến b ```i=&b``` và chạy lại_

```golang
package main

import "fmt"

type Book struct {
	author, title string
}

type Magazine struct {
	title string
	issue int
}

func (b *Book) Assign(n, t string) {
	b.author = n
	b.title = t
}
func (b *Book) Print() {
	fmt.Printf("Author: %s, Title: %s\n", b.author, b.title)
}

func (m Magazine) Assign(t string, i int) {
	m.title = t
	m.issue = i
}
func (m Magazine) Print() {
	fmt.Printf("Title: %s, Issue: %d\n", m.title, m.issue)
}

type Printer interface {
	Print()
}

func main() {
	var b Book                                 // Declare instance of Book
	var m Magazine                             // Declare instance of Magazine
	b.Assign("Jack Rabbit", "Book of Rabbits") // Assign values to b via method

	var i Printer // Declare variable of interface type
	fmt.Println("Call interface")
	i = &b    // Method has pointer receiver
	i.Print() // Show book values via the interface
	i = m    // Magazine also satisfies shower interface
	i.Print() // Show magazine values via the interface
}
```

## Polymorphism (đa hình)

Thay vì implement 1 cách hơi tường minh trong go. Ta sử dụng hàm với parameter là interface thì khi pass vào sẽ thỏa mãn interface.

```golang
package main

import (
	"fmt"
)

// Geometry is an interface that defines Geometrical Calculation
type Geometry interface {
	Edges() int
}

// Pentagon defines a geometrical object
type Pentagon struct{}

// Hexagon defines a geometrical object
type Hexagon struct{}

// Edges implements the Geometry interface
func (p Pentagon) Edges() int { return 5 }

// Edges implements the Geometry interface
func (h Hexagon) Edges() int { return 6 }

// Parameter calculate parameter of object
func Parameter(geo Geometry, value int) int {
	num := geo.Edges()
	calculation := num * value
	return calculation
}

func main() {
	p := new(Pentagon)
	h := new(Hexagon)

	g := [...]Geometry{p, h}

	for _, i := range g {
		fmt.Println(Parameter(i, 5))
	}
}
```

## Interface values

Interface value là cặp tuple value và kiểu cụ thể 

```
(value, type)
```

Hàm in ra value và type của interface

```golang
package main

import (
	"fmt"
	"math"
)

type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	fmt.Println(t.S)
}

type F float64

func (f F) M() {
	fmt.Println(f)
}

func main() {
	var i I

	i = &T{"Hello"}
	describe(i)
	i.M()

	i = F(math.Pi)
	describe(i)
	i.M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
```

```cmd
(&{Hello}, *main.T)
Hello
(3.141592653589793, main.F)
3.141592653589793
```

### Interface với value là nil 

Nếu khai báo struct T là nil value thì khi thực hiện việc gán sang interface thì value của interface sẽ là nil nhưng type sẽ là kiểu implement interface.

```golang
package main

import "fmt"

type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

func main() {
	var i I

	var t *T
	i = t
	describe(i)
	i.M()

	i = &T{"hello"}
	describe(i)
	i.M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
```

Kết quả in ra là

```cmd
(<nil>, *main.T)
<nil>
(&{hello}, *main.T)
hello
```

### Type assertion

Nhắc lại, do interface chứa cặp value và kiểu cụ thể (kiểu cụ thể ko đúng lắm) nên để ép kiểu interface về kiểu cụ thể ta dùng type assertion.

Sử dụng type assertion cho phép lấy lại giá trị của kiểu cụ thể của kiểu implement interface. Sử dụng syntax này sẽ quăng ra ```panic``` nếu assert ko phải là kiểu T.

```golang
t := i.(T)
```

Golang cung cấp cách để test liệu giá trị interface có chứa kiểu cụ thể mà người dùng mong muốn hay ko, sử dụng khai báo như bên dưới

```golang
t, ok := i.(T)
```

Giá trị trả về gồm 2 giá trị: giá trị kiểu cụ thể và boolean để kiểm tra. Nếu ```i``` chứa kiểu T thì ```ok``` sẽ là true và t sử dụng được. Ngược lại, ```ok``` sẽ là false và ```t ``` sẽ là zero value của kiểu ```T``` và panic.

Kiểu ```Pentagon``` bên dưới implement interface ```Polygons``` bằng cách gán biến sau đó dùng type assert để về lại kiểu Pentagon.

**Lưu ý: syntax của type assert là luôn phải có vế trái.**

```golang
package main

import "fmt"

type Polygons interface {
	Perimeter()
}

type Object interface {
	NumberOfSide()
}

type Pentagon int

func (p Pentagon) Perimeter() {
	fmt.Println("Perimeter of Pentagon", 5*p)
}

func (p Pentagon) NumberOfSide() {
	fmt.Println("Pentagon has 5 sides")
}

func main() {
	var p Polygons = Pentagon(50)
	p.Perimeter()
	var o Pentagon = p.(Pentagon)  // type assert
	o.NumberOfSide()

	var obj Object = Pentagon(50)
	obj.NumberOfSide()
	var pent Pentagon = obj.(Pentagon)
	pent.Perimeter()
}
```

### Type Switches

Thay vì dùng trực tiếp kiểu khai báo trên, ta có thể dùng if-else hoặc switch để kiểm tra.

```
switch v := i.(type) {
case T:
    // here v has type T
case S:
    // here v has type S
default:
    // no match; here v has the same type as i
}
```

### Type assertion interface

Bên trên tôi trình bày cách từ interface ép kiểu về type cụ thể. Nhưng đôi khi bạn muốn assert xem T có implement interface hay ko. Syntax assert vẫn như cũ. Code bên dưới kiểu Pentagon implement 2 interface Polygons và Object. Sau khi gọi method ```p.Perimeter()``` từ ```interface Polygons```, ta ép kiểu về p về ```interface Object```. Để kiểm tra xem có thỏa mãn sử dụng ```t, ok := i.(T)```

**Lưu ý: Khi Type assert kiểu gì thì biến nhận phải là kiểu đó, ở đây là interface object.** Dù khi print ra bằng ```"%T"``` thì vẫn là kiểu cụ thể

_Thử xóa hàm ```NumberOfSide``` rồi chạy lại_

```golang
type Polygons interface {
	Perimeter()
}

type Object interface {
	NumberOfSide()
}

type Pentagon int

func (p Pentagon) Perimeter() {
	fmt.Println("Perimeter of Pentagon", 5*p)
}

func (p Pentagon) NumberOfSide() {
	fmt.Println("Pentagon has 5 sides")
}

func main() {
	var p Polygons = Pentagon(50)
	p.Perimeter()
	var o Object = p.(Object)  // type assert interface
	o.NumberOfSide()
	fmt.Printf("%T", o) 	// Check the concrete type held by the Object interface 
}
```

## Empty interface

### Type assert với empty interface

```golang
type Stringer interface {
    String() string
}

var value interface{} // Value provided by caller.
switch str := value.(type) {
case string:
    return str
case Stringer:
    return str.String()
}
```

## Function type implement interface

Trong ```package net/http``` có interface Handler. Handler dùng để implement web. Ví dụ bên dưới type database sẽ implement ServeHTTP

```golang
type Handler interface {
	ServeHTTP(w http.ResponseWriter, r *http.Request)
}
```

```golang
package main

import (
	"net/http"
	"fmt"
)

type dollars float32

type database map[string]dollars

func (d database) ServeHTTP(resp http.ResponseWriter, requ *http.Request){
	switch req.URL.Path {
	case "/list":
		for item, price := range d{
			fmt.Fprintf(resp, "%s: %s\n", item, price)
		}
	default:
		resp.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(resp, "no such page: %s\n", requ.URL)

		// Equivalently
		// msg := fmt.Sprintf("no such page: %s\n", requ.URL)
		// http.Error(resp, msg, http.StatusNotFound)
	}
}

func main(){
	db := database{"shoes" : 50, "socks": 4}
	http.ListenAndServe("localhost:8080", db)
}
```

Tuy nhiên việc viết ServeHTTP ko có tính linh hoạt do type database phải implement ServeHTTP thì mới chạy đc web, nên mỗi khi cần thêm path hoặc chỉnh sửa logic thì phải sửa lại ServeHTTP.
Do vậy package golang định nghĩa ra type là function HandlerFunc và function này sẽ implement ServeHTTP, khi sử dụng chỉ cần cast nó lại thành function HandlerFunc miễn là prototype hàm giống nhau là có thể cast được. Điều này cho phép người lặp trình tự do định nghĩa logic và ```type database``` ko còn kết dính vào ServeHTTP

```golang
type HandlerFunc func(http.ResponseWriter, *http.Request)

func (f HandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f(w, r)
}
```

```golang 

package main

import (
	"net/http"
	"fmt"
	"log"
)

type dollars float32

type database map[string]dollars

func(db database) listAll(w http.ResponseWriter, r *http.Request){
	for item, price := range d{
			fmt.Fprintf(resp, "%s: %s\n", item, price)
	}
}

func main(){
	db := database{"shoes" : 50, "socks": 4}
	mux := http.NewServeMux()
	mux.Handle("/list", http.HandlerFunc(db.listAll))	// cast db.listAll sang http.HandlerFunc
	log.Fatal(http.ListenAndServe("localhost:8080", db))
}
```

## Embedding an interface within a struct

Khi nhúng interface vào trong struct thì các phương thức của interface được promoted thành phương thức của struct Job, và điều này sẽ làm cho struct này như đã implement interface Worker. Thực tế thì vẫn cần 1 struct khác - Developer - để implement Worker

https://eli.thegreenplace.net/2020/embedding-in-go-part-3-interfaces-in-structs/

```golang
type Worker interface {
    Work() string
    Rest() string
}

// Job implement all Woker's method
type Job struct {
    Worker
}

// Implement the interface with a concrete type
type Developer struct{}

func (d Developer) Work() string {
    return "Developing software"
}

func (d Developer) Rest() string {
	return "end of day -- sleep"
}

// Use the struct with the embedded interface
func main() {
    // Assign an instance of Developer to the Worker field of Job
    j := Job{Worker: Developer{}}
    
    // Call the method of the embedded interface
    result := j.Work()
    fmt.Println(j.Work()) // Outputs: Developing software
    fmt.Println(j.Rest()) // Outputs: end of day -- sleep
}
```

Khi nhúng như vậy, struct Job có thể chỉ cần override 1 trong các phương thức của interface. Các phương thức còn lại giữ nguyên, giúp tái sử dụng dễ dàng.
Nếu cần decorate thì Rest của job gọi ```j.Worker.Rest()``` để gọi phần impl của ```struct Developer```

```golang
package main

import "fmt"

type Worker interface {
    Work() string
    Rest() string
}

// Job implement all Woker's method
type Job struct {
    Worker
}

// override Rest()
func (j Job) Rest() string{
	return "over time 2 hours " // + j.Worker.Rest()
}

// Implement the interface with a concrete type
type Developer struct{}

func (d Developer) Work() string {
    return "Developing software"
}

func (d Developer) Rest() string {
	return "end of day -- sleep"
}

// Use the struct with the embedded interface
func main() {
    // Assign an instance of Developer to the Worker field of Job
    j := Job{Worker: Developer{}}
    
    // Call the method of the embedded interface
    fmt.Println(j.Work()) // Outputs: Developing software
    fmt.Println(j.Rest()) // over time 2 hours
}
```

Ví dụ gói net có interface Conn với 8 phương thức cần implment. Tuy nhiên đôi khi bạn chỉ muốn override 1 vài phương thức nhưng nếu viết đầy đủ hết 8 phương thức này ko thực sự cần thiết. Sử dụng Embedding interface trong struct giúp giải quyết tình huống như vậy

```golang
type StatsConn struct {
  net.Conn

  BytesRead uint64
}

func (s *StatsConn) Close() error {
	BytesRead = 0
	return s.Conn.Close()
}
``` 

### Ví dụ từ package sort

Trong package này để sort. sử dụng ```sort.Sort```. Hàm Sort sẽ sử dụng các method của interface Interface mà đã được implement để sort. Package sort cũng đã cung cấp type ```type IntSlice []int``` đã implement Interface

```golang
type Interface interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}
```

```golang
package main

import (
	"fmt"
	"sort"
)

func main() {
	s := []int{5, 2, 6, 3, 1, 4} // unsorted
	sort.Sort(sort.IntSlice(s))
	fmt.Println(s)  // output: [1 2 3 4 5 6]
}
```

Ở đây chưa có gì đặc biệt. Điều đặc biệt là package này cung cấp ```type reverse struct {}``` dùng để sort ngược. Type này nhúng interface Interface và chỉ override mỗi phương thức ```Less(i, i int)```. Hàm Reverse sẽ new type reverse, hàm nhận vào interface và trả ra cũng là 1 interface 

```golang
type reverse struct {
  sort.Interface
}

func (r reverse) Less(i, j int) bool {
  return r.Interface.Less(j, i)
}

func Reverse(data sort.Interface) sort.Interface {
  return &reverse{data}
}
```

Ở hàm main để sort ngược. Ta cần gọi sort.Reverse để tạo struct reverse. Tuy nhiên vì struct reverse chỉ override mỗi Less nên ta cần phải cung cấp type thực sự đã implement các hàm của Interface, cụ thể là sort.IntSlice type này đã implement các phương thức khác. sort.Sort sẽ thực hiện việc sort

```golang
package main

import (
	"fmt"
	"sort"
)

func main(){
	s := []int{5, 2, 6, 3, 1, 4} // unsorted
	sort.Sort(sort.Reverse(sort.IntSlice(s)))
	fmt.Println(s)  // output: [6 5 4 3 2 1]
}
```

### context.WithValue

Package context có hàm WithValue, nó trả về copy của parent và key và value của parent.

Code đã bỏ qua checking lỗi 
```
func WithValue(parent Context, key, val interface{}) Context {
  return &valueCtx{parent, key, val}
}
```

valueCtx chỉ cần override Value

```
type valueCtx struct {
  Context
  key, val interface{}
}

func (c *valueCtx) Value(key interface{}) interface{} {
  if c.key == key {
    return c.val
  }
  return c.Context.Value(key)
}
```

### os.File, io.ReaderFrom

Trong os.File type File implement ReaderFrom. Với f.readFrom(r) sẽ gọi syscall của hệ điều hành thay vì là hàm tự viết, điều này cung cấp cho việc copy nhanh vì thực hiện trực tiếp trong kernel. Method này sẽ trả về false nếu ko gọi đc syscall, nếu thất bại sẽ gọi hàm genericReadFrom

```golang
type ReaderFrom interface {
    ReadFrom(r Reader) (n int64, err error)
}
```
```golang
func (f *File) ReadFrom(r io.Reader) (n int64, err error) {
  if err := f.checkValid("write"); err != nil {
    return 0, err
  }
  n, handled, e := f.readFrom(r)
  if !handled {
    return genericReadFrom(f, r)
  }
  return n, f.wrapErr("write", e)
}
```

Hàm genericReadFrom sẽ gọi io.Copy và file cần phải wrap vào type là onlyWriter. Tuy nhiên câu hỏi là tại sao cần phải wrap lại như vậy?

```golang
func genericReadFrom(f *File, r io.Reader) (int64, error) {
  return io.Copy(onlyWriter{f}, r)
}
```

```golang
type onlyWriter struct {
  io.Writer
}
```

Struct onlyWriter sử dụng interface embedded trong struct, **onlyWriter chỉ implement interface io.Writer**. Tức type File mới là nơi đã implement interface Writer. Tác giả đã check quanh onlyWriter và nhận thấy là onlyWriter ko override bất kỳ phương thức nào. Vậy tại sao vẫn cần việc wrap này?

Hàm io.Copy sẽ gọi copyBuffer (vì đây hàm khá dài nên lược bỏ bớt). Điều thú vị là trong những dòng đầu của ```copyBuffer```, để ý mà xem, ```dst``` tức ```onlyWriter``` sẽ check type assert interface ReaderFrom để có thể ưu tiên gọi hàm syscall thay vì hàm tự viết. Nếu ở đây ta pass là ```type File``` thay vì là wrap ```type onlyWriter``` thì sẽ gọi ReadFrom vô tận. Bằng các wrap này ta sẽ thực hiện đc các công đoạn bên dưới, do onlyWriter chỉ implement interface io.Writer nên type assert interface ReaderFrom sẽ failed dù rằng ```type File có method ReadFrom```.

```golang
func copyBuffer(dst Writer, src Reader, buf []byte) (written int64, err error) {
	// If the reader has a WriteTo method, use it to do the copy.
	// Avoids an allocation and a copy.
	if wt, ok := src.(WriterTo); ok {
		return wt.WriteTo(dst)
	}
	// Similarly, if the writer has a ReadFrom method, use it to do the copy.
	if rt, ok := dst.(ReaderFrom); ok {
		return rt.ReadFrom(src)
	}


	// do write ....
	// ...
	// ...
}
```

Việc định nghĩa type onlyWriter giúp ta dễ hiểu chuyện gì đang xảy ra. Dù vậy ở ```package tar``` tác giả viết package này sử dụng anonymous struct
```golang
io.Copy(struct{ io.Writer }{sw}, r)
```


---

## Interface nâng cao

Sau khi tìm hiểu qua interface cơ bản cách implement và type assert thì sau đây là cách sử dụng interface trong project. 

## Phế võ công

Trước khi tiếp tục ta cần phế võ công về interface đã được học với các ngôn ngữ hướng đối tượng khác như java, c#, ...

Việc phế này ko phải toàn bộ: 

Các thứ cần phế: Tính kế thừa, khai báo interface là duy nhất trong project, sử dụng từ khóa "implements", ưu tiên khai báo interface trước như sự trừu tượng và để cho các lớp con định nghĩa nó tức định nghĩa interface trước, hàm trả về là interface. 

hàm trả về là interface. ví dụ Inter là interface, subclass implemet Inter

```java
public Inter a() {
	return Inter = new Subclass();
}
```

## Go Interface

Trước khi bắt cần hiểu vài thứ

Đầu tiên cần hiểu interface như là collection các định nghĩa phương thức trong package. (1)

Thứ 2: interface nên càng ngắn các phương thức càng tốt, vì càng ngắn thì dễ dàng thỏa mãn interface. (2)

Thứ 3: Interface nên được khai báo trong trong package sẽ sử dụng. _Abtraction phải khám phá, ko phải định nghĩa trước_  (3)

Thứ 4: Khẩu quyết _Nhận vào interface, trả về 1 struct hoặc type cụ thể_ (4)

Thứ 5: Phương thức ở các package có thể giống nhau nhưng 1 struct chỉ thực sự implement tại thời điểm coding. Ví dụ: Pentagon chỉ implement khi thực hiện phép gán hoặc pass qua parameter. (xem lại type assertion để check type nếu ko thỏa interface).

```golang

type Polygons interface {
	Perimeter()
}

type Pentagon int

func (p Pentagon) Perimeter() {
	fmt.Println("Perimeter of Pentagon", 5*p)
}


func main() {
	var p Polygons = Pentagon(50) // implement
	p.Perimeter()
}

```

Trước khi đi sau vào cần làm rõ vài ý về nơi interface sẽ khai báo. Có 2 cách mà interface đc khai báo.

- Interface được khai báo trong cùng package (hoặc sub package của nó) mà sẽ implement. Ví dụ: package a_pack khai báo interface và đồng thời trong package này cũng là nơi implement.
- Interface khai báo ở bên ngoài package, nơi implement sẽ ở 1 package khác. Ví dụ package a_pack khai báo interface, package b_pack là nơi implement interface của package a_pack. Điều (3) chính xác là cách này.

Trong Go được khuyến khích sử dụng cách 2. Nhưng cũng có ngoại lệ như ```package encoding/json``` là nơi khai báo interface, nhưng cũng đồng thời implement

https://pkg.go.dev/encoding/json#Marshaler


### Ví dụ

Ví dụ này sẽ đi theo điều (1), (2), và (3). 

Viết hàm in giả lập như ```package fmt```. Với package cmdout này tôi mong muốn bất cứ phương thức nào có khai báo giống như phương thức trong ```interface Stringer``` thì sẽ thỏa mãn và xài được hàm Print. Trong hàm Print sẽ sử dụng fmt.Printf. 

Như ta đã biết hàm ```fmt.Printf``` nhận vào 1 string. Nên ta cần tạo 1 interface Stringer khi pass vào hàm nếu giá trị pass vào có hàm ```String() string``` thì ta sẽ gọi nó. Lưu ý đoạn assert Stringer. 

package cmdout

```golang
package cmdout

import "fmt"

type Stringer interface{
	String() string
}

Print(a ...interface{}) (error){
	if a == nil {
		return fmt.Errorf("the Print's parameter is nil")
	}
	switch str := value.(type) {
	case error:
		return str.Error()
	case string:
	    fmt.Printf(str)
	case Stringer:
	    fmt.Printf(a.String())
	default:
		fmt.Printf("type %T not implements Stringer", str)
	}
	return nil
}
```

Như đã trình bày bên trên thì interface Stringer như là collection các hàm mà package cmdout này sử dụng miễn là các struct có phương thức thỏa là sẽ dùng được hàm Print. Cụ thể là Person đã thỏa mãn interface Stringer.

Phương thức PrintPerson của Person sử dụng hàm Print của package cmdout

package employee

```golang
package employee

type Person struct {
	Name string
	Age  int
}

func (p Person) String () string{
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func (p Person) PrintPerson(){
	cmdout.Print(p)
}
```

trong main

```golang
package main

import (
	"github/employee"
	"github/cmdout"
)

func main(){
	person := employee.Person{"Arthur Dent", 42}
	cmdout.Print(person)
	person.PrintPerson()
}
```

## Phân tích, ưu, khuyết điểm của interface

### So sánh với ngôn ngữ khác

Ở ví dụ cmdout và employee. Thay vì package Person phải khai báo interface rồi implement những hàm đó trong package Person như các ngôn ngữ khác, thì trong Go package sử dụng mới là nơi phải khai báo interface và interface sẽ trú ngụ trong package sử dụng. Với package Person chỉ viết như là hàm bình thường.

### So với java 

package github.cmdout. Tên hàm đổi thành Print do String trùng từ khóa trong java. Trong ngôn ngữ java thì cần khai báo interface như là duy nhất trong project (khác package ... ), interface phải đc định nghĩa nghĩa trước

```java
package github.cmdout

public interface Stringer {
	String Print();
}

```

Chú ý: Package implement thường quy ước đc đặt ở vị trí sub-package của cmdout. Tuy nhiên, bạn có thể đặt ở vị trí khác như package github.employee

package github.cmdout.employee

```java
package github.cmdout.employee

public class Employee implements Stringer {
	
	String name;
	int age;

	@Override
	public String Print(){
		return String.format("%s (%d years old)", name, age);
	}
}
```

package service. Sử dụng interface Stringer, hàm PrintSerice tương đương với hàm Print trong package cmdout trong Go

```java
package github.service

import github.cmdout;
import github.employee;

public class ServiceEmployee {
	public void PrintService(){
		Stringer s = new Employee();
		System.out.println(s.Print());
	}
}
```

Như ta thấy trong java, ServiceEmployee trong package service. Mong đợi sử dụng interface Stringer và các lớp con implement của nó. Các lớp con phải sử dụng từ khóa "implements" và chỉ định đúng interface sẽ implement ví dụ ở đây là Stringer.

Còn trong Go việc tạo interface sẽ được chuyển sang cho package sẽ dùng tới interface Stringer này. Các phương thức mà implement interface sẽ cần biết hoặc ko cần biết đến sự tồn tại của interface Stringer. (Khá giống với mẫu adapter). 

Như Person trong Go. Phương thức String thực tế ko cần biết đến interface Stringer, khác với java là person phải cần biết để có thể implement được interface. Phương thức của Person viết như hàm bình thường. Và việc của package cmdout muốn có được sự implementation sẽ khai báo interface rồi sử dụng interface này.

Do vậy lời khuyên của ngôn ngữ go là interface nên càng ngắn càng tốt (2) và (3), vì càng ngắn thì các struct dễ thỏa mãn interface hơn.

### Ưu điểm

Vì interface như là tập hợp nên ở package nào sử dụng chỉ cần định nghĩa các interface thì sẽ thỏa mãn.

### Khuyết điểm 

Việc này có thể dẫn tới phải khai báo lặp lại các interface trong các package khác nhau. ~~ví dụ trong ```package io.fs```. File là interface bao gồm các phương thức Stat, Read, Close.???~~

```golang
type File interface {
	Stat() (FileInfo, error)
	Read([]byte) (int, error)
	Close() error
}
```

## Sử dụng type assertion để truy vấn hành vi

Hàm writeHeader này sẽ chuyển chuỗi "Conetent-Type: ..." thành []byte rồi ghi header của HTTP response.

```golang
func writeHeader(w io.Writer, contentType string) error {
	if _, err := w.Write([]byte("Content-Type: ")); err != nil {
		return err
	}
	if _, err := w.Write([]byte(contentType)); err != nil {
		return err
	}
	// ...
}
```

Việc chuyển đổi từ string sang []byte cần cấp phát bộ nhớ và tạo bản copy (pass by value), bản copy này sau khi ghi xong sẽ thu hồi bởi GC. Tuy nhiên việc cấp phát bộ nhớ sẽ làm chậm web server.

Giải pháp là sử dụng interface định nghĩa WriteString() và assert. Vì nếu ```w``` có method WriteString thì việc ghi hiệu quả hơn, tránh việc cấp phát

```golang
type stringWriter interface {
	WriteString(string) (n int, err error)
}
```

Hàm writeString khai báo ```interface stringWriter``` với hàm là ```WriteString```. Sử dụng assert để thăm dò kiểu cụ thể mà interface io.Writer đang giữ có hàm ```WriteString``` hay ko. Nếu có thì sử dụng nó, nếu ko thì cấp phát bộ nhớ tạm.

Kỹ thuật ở đây là interface được khai báo ở nơi sẽ sử dụng và sử dụng assert để "cast" thành interface stringWriter. Đối với ngôn ngữ như java, c# ta dùng intancesof

```golang
func writeString(w io.Writer, s string) (n int, err error) {
	type stringWriter interface {
		WriteString(string) (n int, err error)
	}
	if sw, ok := w.(stringWriter); ok {
		return sw.WriteString(s) // avoid a copy
	}
	return w.Write([]byte(s)) // allocate temporary copy
}

func writeHeader(w io.Writer, contentType string) error {
	if _, err := writeString(w, "Content-Type: "); err != nil {
		return err
	}
	if _, err := writeString(w, contentType); err != nil {
		return err
	}
	// ...
}
```

Đối với java để check B có implement interface hay ko ta dùng instanceof. Tuy nhiên, vì class B là khai báo tường minh nên ko có sự linh động để như Go.

Ví dụ: có 2 interface I và II. Class B implements 2 interface này, class C sẽ implement interface I. Class A sẽ kiểm tra xem class B và C có implement 2 interface này tại thời điểm run-time. 

```java
public interface I { 
	void i(); 
}

public interface II { 
	void ii (); 
}

public class C implements I {
	public void i (){
		System.out.println("c - i");
	}
}

public class B implements I, II{
	public void ii(){
		System.out.println("b - ii"); 
 	}
  
  	public void i(){
    	System.out.println("b - i");
  	}
}

public class A {
	public void func(I i){
   		if (i instanceof II x){
     		x.ii();
   		}
   
   		if (i instanceof I x){
    		x.i(); 
   		}
 	}

 	public static void main (String []args){	 	
		A a = new A();
		a.func(new B());
		a.func(new C());
 	}
}

```

## Nhận vào interface, trả về 1 struct hoặc type cụ thể

Khác với ngôn ngữ **như java mong đợi** nhận vào class cụ thể hoặc interface và **trả về là interface**. Hoặc mong đợi chỉ trả về là interface. Sau đó khi sử dụng sẽ gọi phương thức mà interface định nghĩa.

Đối với Go khuyến khích _Nhận vào interface, trả về 1 struct hoặc type cụ thể_. Tuy nhiên điều này ko phải lúc nào cũng đúng 100% ở việc trả về là 1 struct hoặc type cụ thể. ví dụ như hàm ```sort.Reverse``` ở ví dụ sort [Embedding an interface within a struct](#Embedding-an-interface-within-a-struct)

### So với java

Ưu tiên trả về interface

```java
public interface List{
	void func();
}

public List a(Subclass sub){
	return sub; // subclass implement List
}

public List b(){
	return new Subclass(); // subclass implement List
}

// gọi hàm func() định nghĩa bởi interface
public boolean process(List inter){
	inter.func();

	// hoặc
	List list = new Subclass();
	list.func();
}

```

Trong Go, hàm mong đợi trả về 1 struct cụ thể thay vì là interface. ```struct``` này có thể là implements interface hoặc ko implement interface. Ở đây tôi ví dụ là struct trả về sẽ implement interface.

Hãy xem ví dụ về gzip và file. ```io.Reader``` là 1 interface trong ```package io```

```golang
type Reader interface {
	Read(p []byte) (n int, err error)
}
```

package os trả về con trỏ file. ```struct file``` này có phương thức Read() thỏa mãn ```interface io.Reader```. Hàm ```NewReader``` của gzip nhận vào ```interface io.Reader``` và trả về con trỏ kiểu ```Reader``` (lưu ý kiểu Reader là kiểu cụ thể ko phải là interface).

Hàm process nhận vào ```interface io.Reader```. Vì ta biết ```Reader``` của ```package gzip``` đã implements ```io.Reader``` nên khi pass vào sẽ thỏa mãn. 

Để sử dụng hàm Read() ta sử dụng trực tiếp từ gzip bằng cách gọi ```zip.Read()``` hoặc sử dụng nó thông qua các hàm của ```package io```, nơi mà sẽ sử dụng như ví dụ trước. Trong ```io.Copy``` sẽ sử dụng method Read().

```golang
package main

import (
	"os"
	"gzip"
	"io"
)

// In ra console
func process(r io.Reader) error{
	io.copy(os.Stdout, r)
	return nil
}

func main(){
	file, _ := os.Open("")
	defer file.Close()
	zip, _ := gzip.NewReader(file)

	// gọi zip.Read()
	// b := make([]byte, 8)
	// zip.Read(b)

	defer zip.Close()
	process(zip)
}
```

Bên trong ```package gzip```. Phương thức ```z.Reset()``` chịu trách nhiệm tạo các property của Reader, NewReader cấp phát vùng nhớ.

```go
type Reader struct {
	// chi tiết xem trong package gzip
}

func NewReader(r io.Reader) (*Reader, error) {
	z := new(Reader)
	if err := z.Reset(r); err != nil {
		return nil, err
	}
	return z, nil
}

func (z *Reader) Reset(r io.Reader) error {
	*z = Reader{
		decompressor: z.decompressor,
		multistream:  true,
	}
	if rr, ok := r.(flate.Reader); ok {
		z.r = rr
	} else {
		z.r = bufio.NewReader(r)
	}
	z.Header, z.err = z.readHeader()
	return z.err
}

func (z *Reader) Read(b []byte) (n int, err error){
	// implement , xem trong package gzip
}
```

### Ví dụ 

Trong ví dụ này ta sẽ implement giống cách mà package gzip làm nhưng là parser file csv. Cũng nhận vào file nhưng trả về struct


### Trả về là interface

Không phải lúc nào cũng trả ra là struct. Như trong ```package io```. 

```golang
func LimitReader(r Reader, n int64) Reader{
	return &LimitReader{r, n}
}
```

## Struct có thuộc tính là interface 

Cũng giống như các ngôn ngữ java, c#,... Interface có thể là thuộc tính của một class.

```java
public interface Interf {
	void func();
}

public class Service {
	Interf interf;  // thuộc tính của Service

	public Service(Interf interf){
		this.interf = interf;
	}

	public void invokeService(){
		interf.func();
	}
}

public class App {
	public static void main (string []args){
		Subclass subclass = new Subclass()
		Service service = new Service(subclass);
		// or 
		// Interf interf = new Subclass()
		// Service service = new Service(interf);
		service.invokeService();
	}
}
```

Khám phá ```package bufio``` sử dụng interface Reader. Trong package bufio có struct là ```Scanner``` có thuộc tính là interface io.Reader. Do Go ko có hàm tạo (constructor) nên hàm ```NewScanner``` hoạt động như là hàm tạo. 

Cũng giống như điều (4) của Go, tuy nhiên ở đây ```NewScanner()``` trả về struct ko implement ```interface Reader```. Ở hàm main sử dụng như việc tạo service trong java.

```golang
// trong package io
type Reader interface {
	Read(p []byte) (n int, err error)
}

// trong package bufio
type Scanner struct {
	r 		io.Reader // The reader provided by the client.
	buf 	[]byte
}

func NewScanner(r io.Reader) *Scanner {
	return &Scanner{
		r:  r,
	}
}

func (s *Scanner) Scan() bool {
	// implement, xem chi tiết trong bufio
	// gọi phương thức read ở trong Scan
	// Lưu ý: s.buf cần được cấp phát trước khi pass vào s.r.Read()
	s.r.Read(s.buf)
}

func (s *Scanner) Text() string {
	// implement
}

func main(){
	file, _ := os.Open("file.txt")
	scan := bufio.NewScanner(r)

	for scan.Scan(){
		fmt.Println(scan.Text())
	}
}
```

### Ví dụ: 