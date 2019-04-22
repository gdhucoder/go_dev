package main

import (
	"fmt"
	"os"
	"time"
)

type PathError struct {
	path       string
	op         string
	message    string
	createTime string
}

func (p *PathError) Error() string {
	return fmt.Sprintf("path=%s\nop=%s\nmessage=%s\ncreateTime=%s\n",
		p.path, p.op, p.message, p.createTime)
}

func Open(filename string) (err error) {
	file, err := os.Open(filename)
	if err != nil {
		return &PathError{
			path:       filename,
			op:         "read",
			message:    err.Error(),
			createTime: fmt.Sprintf("%v", time.Now()),
		}
	}
	defer file.Close()
	return nil
}

// 自定义错误要实现Error接口
func main() {
	err := Open("c:llsdfsdf.txt")
	if err != nil {
		fmt.Println(err)
	}
	// path=c:llsdfsdf.txt
	// op=read
	// message=open c:llsdfsdf.txt: The system cannot find the file specified.
	// createTime=2019-04-21 09:56:56.4403007 +0800 CST m=+0.011967601

	v, ok := err.(*PathError)
	if ok {
		fmt.Println(v)
	}

	switch etp := err.(type) {
	case *PathError:
		fmt.Println(err)
	default:
		fmt.Println(etp)
	}
}
