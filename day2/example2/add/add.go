package add

import (
	"fmt"
	// 使用 _ 只调用初始化
	_ "go_dev/day2/example2/test"
)

// step 1声明，初始化
var Name string

var Age int

func Test() {
	fmt.Println("in add ... Test")
	Name = "王天娲"
	Age = 16
}

// step 2 调用 init函数
func init() {
	fmt.Println("in add ... init")
	Name = "胡国栋"
}

// var name = "xxx" 编译的时候就确定值了
