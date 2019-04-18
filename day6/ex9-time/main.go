package main

import (
	"fmt"
	"time"
)

const DateFormat = "2006-01-02"

func testTime() {
	// "2006-01-02 15:04:05.999999999 -0700 MST"
	now := time.Now().Format(DateFormat)
	fmt.Println(now)
	afterTime := time.Now().AddDate(0, 0, 30) //
	after := afterTime.Format(DateFormat)
	fmt.Println(after)

}

func main() {
	testTime()
}
