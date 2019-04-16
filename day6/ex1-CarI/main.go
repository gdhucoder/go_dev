package main

import "fmt"

type CarI interface {
	Run()
	GetName() string
	DiDi()
}

type TestI interface {
	Hello()
}

type BMW struct {
	Name string
}

func (c *BMW) Run() {
	fmt.Printf("%s is running.\n", c.Name)
}

func (c *BMW) GetName() string {
	return c.Name
}

func (c *BMW) DiDi() {
	fmt.Printf("%s is making sound :DiDi.\n", c.Name)
}

func (c *BMW) Hello() {
	fmt.Printf("%s is saying hello!\n", c.Name)
}

type BYD struct {
	Name string
}

func (c *BYD) Run() {
	fmt.Printf("%s is running.\n", c.Name)
}

func (c *BYD) GetName() string {
	return c.Name
}

func (c *BYD) DiDi() {
	fmt.Printf("%s is making sound :DiDi .\n", c.Name)
}

func main() {
	var car CarI
	fmt.Println(car) // car 没有实现接口

	var bmw = &BMW{
		Name: "BMW-00001",
	}
	car = bmw
	car.Run()
	fmt.Println(car.GetName())
	car.DiDi()

	// BMW-00001 is running.
	// BMW-00001
	// BMW-00001 is making sound :DiDi

	var byd = &BYD{
		Name: "BYD-00002",
	}

	car = byd
	car.Run()
	// BYD-00002 is running.

	var test TestI

	test = bmw   // BMW 实现了两个接口
	test.Hello() // BMW-00001 is saying hello!

}
