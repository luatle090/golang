## Go method

### Khai báo struct

type MyStruct struct{
	Counter int
}

### Constructor 

```
func NewMyStruct() MyStruct{
	return MyStruct{Counter: 0}
}
```

### Khai báo phương thức

```
func (m MyStruct) GetCount() int {
	return m.Counter
}
```

_Phương thức gọi phương thức_

cả hai đều là phương thức của MyStruct nhưng Increase là API public, còn increase là private. Để public gọi được private thì phải gọi qua phương thức thay trực tiếp như java, c#

```Go
func (m MyStruct) Increase(){
	m.increase()		// ko thể gọi trực tiếp, phải gọi thông qua phương thức
}

func (m *MyStruct) increase(){
	m.counter = m.counter + 1 
}
```

```Java
public class A{
	int counter;

	public void increaseCounter(){
		increase();   // gọi trực tiếp vì trong cùng 1 class
	}

	private void increase(){
		increase++;
	}
}

```