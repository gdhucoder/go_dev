package main

import "fmt"

type Reader interface {
	Read()
}

type Writer interface {
	Write()
}

type ReaderWriter interface {
	Reader
	Writer
}

type File struct {
}

func (f *File) Read() {
	fmt.Println("read data")
}

func (f *File) Write() {
	fmt.Println("write data")
}

func Test(rw ReaderWriter) {
	rw.Write()
	rw.Read()
}

func main() {
	var file = new(File)
	var rw ReaderWriter
	rw = file
	rw.Read()
	rw.Write()
	// read data
	// write data

	Test(file)
}
