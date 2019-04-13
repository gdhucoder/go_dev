package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

var lock sync.Mutex
var rwLock sync.RWMutex

// 由于资源具有独占性，同时写一个数据会出现资源竞争。
// 例如我们坐火车要去厕所，厕所要锁门
func testMap() {
	var a map[string]string
	a = make(map[string]string, 5)
	a["a"] = "xxxxxxxxxxxxxx"

	for i := 0; i < 2; i++ {
		go func(b map[string]string) {
			lock.Lock() // 互斥锁
			defer lock.Unlock()
			b["a"] = "sdfsdfsdf"
		}(a)
	}
	lock.Lock()
	fmt.Println(a)
	lock.Unlock()
}

// 读写锁
func testRWMap() {
	var a map[string]int
	a = make(map[string]int, 5)
	a["a"] = 123
	a["b"] = 321
	a["c"] = 45

	var count int32
	count = 0

	// write
	for i := 0; i < 2; i++ {
		go func(b map[string]int) {
			// rwLock.Lock()
			lock.Lock()
			b["a"] = rand.Intn(99999999999999)
			time.Sleep(time.Millisecond * 10)
			// rwLock.Unlock()
			lock.Unlock()
		}(a)
	}

	// read
	for i := 0; i < 100; i++ {
		go func(b map[string]int) {
			for {
				// rwLock.RLock()
				lock.Lock()
				// fmt.Println(b)
				time.Sleep(time.Millisecond)
				// rwLock.RUnlock()
				lock.Unlock()
				atomic.AddInt32(&count, 1) // 原子操作
			}
		}(a)
	}
	// fmt.Println(count) // data
	time.Sleep(time.Second * 2)
	// 用读写锁效率： 109976 次， 读请求多的时候效率相差近100倍！
	fmt.Printf("使用读写锁效率: %d 次\n", atomic.LoadInt32(&count)) //
	// 使用互斥锁效率: 1164 次
	fmt.Printf("使用互斥锁效率: %d 次\n", atomic.LoadInt32(&count)) //

}

func main() {

	// --race检查资源竞争
	// go build -race  D:\project\src\go_dev\day4\ex12-lock

	// PS D:\project> go build --race  D:\project\src\go_dev\day4\ex12-lock
	// PS D:\project> D:\project\ex12-lock.exe
	// ==================
	// WARNING: DATA RACE
	// Write at 0x00c00005c2d0 by goroutine 7:
	//   runtime.mapassign_faststr()
	//       C:/Go/src/runtime/map_faststr.go:202 +0x0
	//   main.testMap.func1()
	//       D:/project/src/go_dev/day4/ex12-lock/main.go:12 +0x64
	// testMap()
	testRWMap()
}
