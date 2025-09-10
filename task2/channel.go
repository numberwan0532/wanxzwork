package main

import (
	"fmt"
)

func snedChan(ch chan<- int) {
	for i := 1; i <= 10; i++ {
		ch <- i
		fmt.Println("send chan value:", i)
	}
	close(ch)
}

func readChan(ch <-chan int) {
	for v := range ch {
		fmt.Println("read chan value:", v)
	}
}

func snedChan1(ch chan<- int) {
	for i := 1; i <= 100; i++ {
		ch <- i
		fmt.Println("send chan value:", i)
	}
	close(ch)
}

func readChan1(ch <-chan int) {
	for v := range ch {
		fmt.Println("read chan value:", v)
	}
}
