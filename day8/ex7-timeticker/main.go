package main

import (
	"fmt"
	"time"
)

func main() {
	// 定时器
	t := time.NewTicker(time.Second)
	for v := range t.C {
		fmt.Println(v)
	}

	t.Stop() // 这么用
	// 	2019-04-24 15:17:24.57087 +0800 CST m=+5.011559101
	// 2019-04-24 15:17:29.5708731 +0800 CST m=+10.011562201
	// 2019-04-24 15:17:34.5707072 +0800 CST m=+15.011396301
}
