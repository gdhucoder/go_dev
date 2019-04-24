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
	// --- FAIL: TestLoad (0.00s)
	// d:\project\src\go_dev\day8\ex9-test\stu_test.go:24: test failed open
	// ./stu11.dat: The system cannot find the file specified.
}
