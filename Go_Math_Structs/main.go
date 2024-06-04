package main

import (
	"fmt"
	"math"
)

func main() {

	m := MyClass{}

	fmt.Println("enter your choice ")
	fmt.Println("1: circle")
	fmt.Println("2: square")
	fmt.Println("3: rectangle")
	fmt.Println("4: triangle")

	var choice int
	fmt.Scan(&choice)

	switch choice {
	case 1:
		m.circle()
	case 2:
		m.square()
	case 3:
		m.rectangle()
	case 4:
		m.triangle()

	}

}

type Area struct{}

func (a Area) circle(radious float32) float32 {
	return math.Pi * radious * radious
}

func (a Area) square(side float32) float32 {
	return side * side
}

func (a Area) rectangle(length, width float32) float32 {
	return length * width
}
func (a Area) triangle(base, height float32) float32 {
	return 0.5 * base * height
}

type MyClass struct {
	Area
}

func (m MyClass) circle() {
	var radious float32
	fmt.Println("enter the radious")
	fmt.Scan(&radious)
	answer := m.Area.circle(radious)
	fmt.Printf("area of the circle is %.2f", answer)
}
func (m MyClass) square() {
	var side float32
	fmt.Println("enter the length of side")
	fmt.Scan(&side)
	answer := m.Area.square(side)
	fmt.Printf("area of the square is %.2f", answer)
}
func (m MyClass) rectangle() {
	var length, width float32
	fmt.Println("enter the length and width")
	fmt.Scan(&length, &width)
	answer := m.Area.rectangle(length, width)
	fmt.Printf("area of the rectangle is %.2f", answer)
}
func (m MyClass) triangle() {
	var base, height float32
	fmt.Println("enter the base and height")
	fmt.Scan(&base, &height)
	answer := m.Area.triangle(base, height)
	fmt.Printf("area of the triangle is %.2f", answer)
}
