package main

import (
	"fmt"
	"sync"
	"time"
)

func userGo() {

	go func() {
		for i := 1; i < 10; i += 2 {
			fmt.Println("奇数:", i)
		}
	}()

	go func() {
		for i := 2; i <= 10; i += 2 {
			fmt.Println("偶数:", i)
		}
	}()
}

type Task struct {
	Name string
	Func func()
}

func ScheduleTasks(tasks []Task) {
	var wg sync.WaitGroup
	wg.Add(len(tasks))

	results := make(chan string, len(tasks))
	for _, task := range tasks {
		go func(task Task) {
			defer wg.Done()
			startTime := time.Now()
			task.Func()
			userTime := time.Since(startTime)
			results <- fmt.Sprintf("任务:%s,耗时:%s", task.Name, userTime)
		}(task)
	}
	wg.Wait()
	close(results)

	for result := range results {
		fmt.Println(result)
	}
}
