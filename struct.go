package main

import ("fmt")

type Student struct {
	name string
	rollNo int
	marks []int
	grades map[string]int
}

type Circle struct {
	x int 
	y int
	radius float64
	area float64
}

func calcArea(c *Circle) {
	const PI float64 = 3.14
	area := (PI * c.radius * c.radius)
	(*c).area = area
}

func (c *Circle) calArea() {
	c.area = 3.14 * c.radius * c.radius
}

func main() {
	fmt.Println("Hello")
	st := Student{
		name: "Abhradip Paul",
		rollNo: 21,
	}
	fmt.Printf("%+v \n", st)
	c := Circle{x: 5, y: 5, radius: 5, area: 0}
	fmt.Printf("%+v", c)
	calcArea(&c)
	fmt.Printf("%+v", c)
}