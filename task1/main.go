package main

import (
	"fmt"
)

func main() {
	fmt.Println("start...")
	nums := []int{4, 1, 2, 1, 2, 5, 4, 5, 6, 6, 7}
	value := findOnleOne(nums)
	fmt.Println("只出现一次的数:", value)

	fmt.Println("是否是回文数:", isPalindrome(1124211))

	fmt.Println("是否是有效的括号:", isValid("(()[]{}{})"))

	fmt.Println("最长公共前缀:", getLongestPrefix([]string{"flower", "fl", "floight"}))

	fmt.Println("数组加一:", getNumsAdd1([]int{9, 5, 9}))
	fmt.Println("数组加一:", getNumsAdd2([]int{9, 5, 9}))

	fmt.Println("删除排序数组中的重复项:", deleteRepetition([]int{1, 1, 2, 2, 3, 4, 4, 4, 4, 5}))

	fmt.Println("合并区间:", mergeIntervals([][]int{{1, 3}, {2, 6}, {8, 16}, {15, 18}, {20, 25}}))

	fmt.Println("两数之和:", findTowNumTotal([]int{2, 7, 11, 15, 18, 39}, 25))
	fmt.Println("两数之和:", findTowNumTotal1([]int{2, 7, 11, 15, 18, 39}, 25))
}
