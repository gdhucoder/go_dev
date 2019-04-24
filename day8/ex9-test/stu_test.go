package main

import (
	"fmt"
	"testing"
)

func TestSave(t *testing.T) {
	var stu = &Student{
		Name: "胡国栋",
		Age:  33,
		Sex:  "man",
	}
	err := stu.Save()
	if err != nil {
		t.Fatalf("test failed %v", err)
	}
}

func TestLoad(t *testing.T) {
	stu := &Student{}
	err := stu.Load()
	if err != nil {
		t.Fatalf("test failed %v", err)
	}
	fmt.Println(stu)
}
