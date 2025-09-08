package main

import (
	"fmt"
)

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
