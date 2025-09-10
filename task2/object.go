package main

import (
	"fmt"
)

type Shape interface {
	Area()
	Perimeter()
}

type Rectangle struct {
}

type Circle struct {
}

func (r Rectangle) Area() {
	fmt.Println("Rectangle Area")
}

func (r Rectangle) Perimeter() {
	fmt.Println("Rectangle Perimeter")
}

func (c Circle) Area() {
	fmt.Println("Circle Area")
}

func (c Circle) Perimeter() {
	fmt.Println("Circle Perimeter")
}

type Person struct {
	Name string
	Age  int
}

type Employee struct {
	Person
	EmployeeID string
}

func (e Employee) PrintInfo() {
	fmt.Printf("Name: %s, Age: %d, EmployeeID: %s\n", e.Name, e.Age, e.EmployeeID)
}
