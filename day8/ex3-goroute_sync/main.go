package main

import (
	"fmt"
	"sync"
	"time"
)

type Task struct {
	N int
}

var (
	m    = make(map[int]uint64) // 全局变量
	lock sync.Mutex             // 互斥锁
)

func calc(t *Task) {
	var sum uint64
	sum = 1
	for i := 1; i <= t.N; i++ {
		sum *= uint64(i)
	}
	lock.Lock()
	m[t.N] = sum
	lock.Unlock()
}

func main() {
	for i := 0; i < 100; i++ {
		t := &Task{N: i}
		go calc(t)
	}

	time.Sleep(5 * time.Second)
	// 使用 -race 判断竞争
	lock.Lock()
	for k, v := range m {
		fmt.Printf("%d! = %v\n", k, v)
	}
	lock.Unlock()

}
