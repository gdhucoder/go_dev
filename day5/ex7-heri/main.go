package main

import "fmt"

type Car struct {
	Name   string
	Weight int
}

type Bike struct {
	Car
	NumofWheel int
}

// 继承，组合
type Train struct {
	c Car
	b Bike
}

func (t *Car) Run() {
	fmt.Println("Car", t.Name, "is Running")
}

func (t *Bike) Run() {
	t.Car.Run()
	fmt.Println("Bike", t.Name, "is Running")
}

func (t *Car) String() string {
	str := fmt.Sprintf("name = %s, weight = %d\n", t.Name, t.Weight)
	return str
}

func main() {
	var bike Bike
	bike.Name = "捷安特"
	bike.Weight = 100
	bike.NumofWheel = 2

	fmt.Println(bike)

	bike.Run() // 调用父类的方法

	fmt.Println(bike)  //
	fmt.Println(&bike) // 取地址
}
