package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("start...")
	num := 10
	p := &num
	fmt.Println("point value before change:", num)
	changePointValue(p)
	fmt.Println("point value after change:", num)

	s := []int{1, 2, 3, 4, 5}
	changeSliceValue(&s)
	fmt.Println("slice value after change:", s)

	userGo()

	time.Sleep(1 * time.Second)

	tasks := []Task{
		{Name: "任务1",
			Func: func() {
				time.Sleep(2 * time.Second)
			}},
		{Name: "任务2", Func: func() {
			time.Sleep(1 * time.Second)
		}},
		{Name: "任务3", Func: func() {
			time.Sleep(3 * time.Second)
		}},
	}
	ScheduleTasks(tasks)

	rectangle := Rectangle{}
	circle := Circle{}

	rectangle.Area()
	rectangle.Perimeter()

	circle.Area()
	circle.Perimeter()

	emp := Employee{
		Person:     Person{Name: "Alice", Age: 30},
		EmployeeID: "E12345",
	}
	emp.PrintInfo()

	ch := make(chan int)
	go snedChan(ch)
	go readChan(ch)

	time.Sleep(1 * time.Second)

	ch1 := make(chan int, 5)
	go snedChan1(ch1)
	go readChan1(ch1)

	time.Sleep(3 * time.Second)

	addTest()

	atomicTest()
}
