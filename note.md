func add(){}

func add(a int, b int) int {}

func add(a int, b int) (int, int){}

注释： // /* */

常量：

const修饰，永远是只读的，不能修改。

包的别名

调用顺序

多个变量初始化

var(
    a int
    d = 9
    e = "hello world"
)


a := 100 这个语句其实是两条语句，等价于

var a int
a = 100

```go
var ss = `adfadsf \n \n` // 原样输出
fmt.Println(ss)

输出
adfadsf \n \n
```


build 代码
go build D:\project\src\go_dev\day3\example4

执行代码
D:\project\example4.exe
