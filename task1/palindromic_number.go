package main

import (
	"fmt"
)

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
