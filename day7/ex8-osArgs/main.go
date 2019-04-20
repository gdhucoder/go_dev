package main

import (
	"fmt"
	"os"
)

func main() {
	// os.Args是string的切片
	fmt.Printf("len of args: %d\n", len(os.Args))
	for i, v := range os.Args {
		fmt.Printf("arg[%d]=%s\n", i, v)
	}
	// arg[0]=C:\Users\x1c\AppData\Local\Temp\go-build446645618\b001\exe\main.exe
	// arg[1]=sdf
	// arg[2]=sdffd
}
