package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Student struct {
	Name string
	Age  int
	Sex  string
}

func (s *Student) Save() (err error) {
	data, err := json.Marshal(s)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = ioutil.WriteFile("./stu.dat", data, 0755) // 写文件
	return
}

func (s *Student) Load() (err error) {
	data, err := ioutil.ReadFile("./stu.dat") // 读文件
	if err != nil {
		fmt.Println(err)
		return
	}

	err = json.Unmarshal(data, s)
	return
}
