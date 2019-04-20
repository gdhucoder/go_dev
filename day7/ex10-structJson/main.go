package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name     string
	NickName string
	Age      int
	Phone    string
}

// json序列化结构体
func main() {
	u := User{
		Name:     "mahuateng",
		NickName: "tod",
		Age:      18,
		// Phone:    "110",
	}

	data, err := json.Marshal(u)

	if err != nil {
		fmt.Printf("json marshal failed, error :%s\n", err)
	}

	fmt.Println(string(data))
	// {"Name":"mahuateng","NickName":"tod","Age":18,"Phone":""}
}
