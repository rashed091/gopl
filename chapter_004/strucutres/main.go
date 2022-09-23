package main

import "fmt"

type Employee struct {
	ID       int
	Name     string
	Address  string
	Salary   float32
	position string
}

type Point struct {
	X, Y int
}

type Circle struct {
	Point
	Radius int
}

type Wheel struct {
	Circle
	Spokes int
}

func main() {
	var emp Employee

	emp.Name = "Rashed"
	emp.ID = 123456
	emp.Address = "Mohammadpur"
	emp.Salary = 158000.0001
	emp.position = "Eng"

	position := &emp.position
	*position = *position + "II"

	fmt.Println(emp)

	var w Wheel
	w.X = 8
	w.Y = 8
	w.Radius = 5
	w.Spokes = 20
	fmt.Println(w)

	z := Wheel{
		Circle: Circle{
			Point:  Point{X: 8, Y: 8},
			Radius: 5,
		},
		Spokes: 20, // NOTE: trailing comma necessary here (and at Radius)
		// if the enclosing brace is in the next line
	}
	fmt.Println(z)
}
