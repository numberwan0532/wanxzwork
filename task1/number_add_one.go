package main

import (
	"fmt"
	"math"
)

// 数组加一
func getNumsAdd1(nums []int) []int {
	n := len(nums)
	var count int
	for i := 0; i < n; i++ {
		count += nums[i] * int(math.Pow(10, float64(n-i-1)))
	}
	count++

	resultNums := []int{}
	n1 := len(fmt.Sprintf("%d", count))
	for i := 0; i < n1; i++ {
		resultNums = append(resultNums, count/int(math.Pow(10, float64(n1-i-1)))%10)
	}
	return resultNums
}

func getNumsAdd2(nums []int) []int {
	n := len(nums)
	for i := n - 1; i >= 0; i-- {
		if nums[i] < 9 {
			nums[i]++
			return nums
		}
		nums[i] = 0
	}
	// 如果最高位进位，需要在前面加1
	return append([]int{1}, nums...)
}
