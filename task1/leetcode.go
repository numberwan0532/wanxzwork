package main

import (
	"fmt"
	"math"
	"sort"
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

// 是否是回文数
func isPalindrome(x int) bool {
	fmt.Println(x)
	if x < 0 {
		return false
	}
	s := fmt.Sprintf("%d", x)
	n := len(s)
	for i := 0; i < n/2; i++ {
		if s[i] != s[n-1-i] {
			return false
		}
	}
	return true
}

// 判断是否是是()[]{}组成的字符串
func isStringValid(s string) bool {
	stack := [6]string{"(", ")", "[", "]", "{", "}"}
	for i := 0; i < len(s); i++ {
		flag := false
		for _, v := range stack {
			if string(s[i]) == v {
				flag = true
				break
			}
		}
		if !flag {
			return false
		}
	}
	return true
}

// 有效的括号
func isValid(s string) bool {
	n := len(s)
	if n%2 != 0 {
		return false
	}
	if !isStringValid(s) {
		return false
	}

	stack := []rune{}
	pairs := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}
	for _, ch := range s {
		switch ch {
		case '(', '[', '{':
			fmt.Println("push:", string(ch))
			stack = append(stack, ch)
			fmt.Println("stack:", string(stack))
		case ')', ']', '}':
			fmt.Println("get:", string(ch))
			fmt.Println("len(stack)-1:", len(stack)-1)
			fmt.Println("pairs[ch]:", string(pairs[ch]))
			fmt.Println("stack[len(stack)-1]:", string(stack[len(stack)-1]))
			if len(stack) == 0 || stack[len(stack)-1] != pairs[ch] {
				return false
			}
			fmt.Println("stack1:", string(stack))
			stack = stack[:len(stack)-1]
			fmt.Println("stack2:", string(stack))
		}
	}
	return len(stack) == 0
}

// 最长前缀
func getLongestPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	byteStrs := []byte{}
	for i := 0; i < len(strs[0]); i++ {
		falg := false
		marr := make(map[string]int)
		for _, v := range strs {
			if i >= len(v) {
				falg = true
				break
			}
			marr[string(v[i])]++
		}
		if falg {
			break
		}
		if len(marr) == 1 {
			for k := range marr {
				byteStrs = append(byteStrs, k[0])
			}
		} else {
			falg = true
			break
		}
		if falg {
			break
		}

	}
	return string(byteStrs)
}

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

// 删除排序数组中的重复项
func deleteRepetition(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	i := 0
	for j := 1; j < len(nums); j++ {
		if nums[j] != nums[i] {
			i++
			nums[i] = nums[j]
		}
	}
	fmt.Println(nums[:i+1])
	return i + 1
}

// 合并区间
func mergeIntervals(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return intervals
	}
	// 先对区间按照起始位置进行排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	merged := [][]int{}
	for _, interval := range intervals {
		// 如果merged为空，或者当前区间的起始位置大于merged中最后一个区间的结束位置，说明没有重叠，直接添加
		if len(merged) == 0 || interval[0] > merged[len(merged)-1][1] {
			merged = append(merged, interval)
		} else {
			// 否则，存在重叠，更新merged中最后一个区间的结束位置
			merged[len(merged)-1][1] = max(merged[len(merged)-1][1], interval[1])
		}
	}
	return merged
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 两数之和
func findTowNumTotal(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{nums[i], nums[j]}
			}
		}
	}
	return []int{}
}

func findTowNumTotal1(nums []int, target int) []int {
	marr := make(map[int]int)
	for i, v := range nums {
		if val, ok := marr[target-v]; ok {
			return []int{nums[val], nums[i]}
		}
		marr[v] = i
	}
	return []int{}
}
