package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type SafeAdd struct {
	mu    sync.Mutex
	value int
}

func (s *SafeAdd) Add() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.value++
}

func (s *SafeAdd) GetValue() int {
	return s.value
}

func addTest() {
	var wg sync.WaitGroup
	safeAdd := SafeAdd{}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				safeAdd.Add()
			}
		}()
	}
	wg.Wait()
	fmt.Println("final value:", safeAdd.GetValue())
}

type Counter struct {
	count atomic.Int64
}

func (c *Counter) Increment() {
	c.count.Add(1)
}
func (c *Counter) GetCount() int64 {
	return c.count.Load()
}

func atomicTest() {
	var wg sync.WaitGroup
	counter := Counter{}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				counter.Increment()
			}
		}()
	}
	wg.Wait()
	fmt.Println("final atomic count:", counter.GetCount())
}
