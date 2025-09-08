package main

import (
	"fmt"
)

// 只出现一次的数
func findOnleOne(nums []int) int {
	fmt.Println(nums)
	marr := make(map[int]int)
	for _, v := range nums {
		marr[v]++
	}
	for k, v := range marr {
		if v == 1 {
			return k
		}
	}
	return 0
}
