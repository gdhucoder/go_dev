package main

import (
	"flag"
	"fmt"
)

func main() {
	var confPath string
	var logLevel int
	flag.StringVar(&confPath, "c", "", "please input conf path")
	flag.IntVar(&logLevel, "l", 1, "please input log level!")
	flag.Parse() // 必须写
	fmt.Printf("confPath :%s\n", confPath)
	fmt.Printf("log level :%d\n", logLevel)
	// PS D:\project> D:\project\ex9-flag.exe -c C:t.log -l 111 -d 3
	// flag provided but not defined: -d
	// Usage of D:\project\ex9-flag.exe:
	//   -c string
	//         please input conf path
	//   -l int
	//         please input log level! (default 1)
}
