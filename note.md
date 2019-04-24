# Go Lang 笔记

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


获取变量的地址，使用&

获取指针类型所指向的值，使用*

var *p int，使用\*p获取p指向的值

var a int = 5
var p \*int = &a

*p = 5 ， p = 0xffff

非常重要：

map slice chan 指针 interface默认以引用的方式传递


https://gobyexample.com/closures

https://en.wikipedia.org/wiki/Anonymous_function

https://en.wikipedia.org/wiki/Closure_(computer_programming)

![2019-04-12-001](https://gitee.com/gdhu/prvpic/raw/master/2019-04-12-001.jpg)

## 关于锁

```go
	// 用读写锁效率： 109976 次， 读请求多的时候效率相差近100倍！
	fmt.Printf("使用读写锁效率: %d 次\n", atomic.LoadInt32(&count)) //
	// 使用互斥锁效率: 1164 次
	fmt.Printf("使用互斥锁效率: %d 次\n", atomic.LoadInt32(&count)) //
```

## 关于interface

空接口没有任何方法，所有类型都实现了空接口

## 关于reflect

运行时反射，可以使程序操作任意类型的对象。静态类型为TypeOf，动态数据为ValueOf

## 关于负载均衡

[负载均衡](https://en.wikipedia.org/wiki/Load_balancing_(computing))

[RoundRobin DNS](https://en.wikipedia.org/wiki/Round-robin_DNS)


## go 写测试案例

go文件以 xxx_test.go结尾

新建方法

cd 到测试代码所在的目录 直接执行go test

如果想看执行输出的日志 go test -v

```go
func TestAdd(t *testing.T) {
	r := add(2, 4)
	if r != 6 {
		t.Fatalf("expecting %d, but %d\n", 6, r)
	}
	t.Logf("test succ")

	// PS D:\project> cd .\src\go_dev\day8
	// PS D:\project\src\go_dev\day8> cd .\ex9-test\
	// PS D:\project\src\go_dev\day8\ex9-test> go test -v
	// === RUN   TestAdd
	// --- PASS: TestAdd (0.00s)
	// 		calc_test.go:10: test succ
	// PASS
	// ok      go_dev/day8/ex9-test    0.506s
	// PS D:\project\src\go_dev\day8\ex9-test>
}
```

失败
```
PS D:\project\src\go_dev\day8\ex9-test> go test -v
=== RUN   TestAdd
--- FAIL: TestAdd (0.00s)
    calc_test.go:8: expecting 6, but 7
FAIL
exit status 1
FAIL    go_dev/day8/ex9-test    0.513s
```