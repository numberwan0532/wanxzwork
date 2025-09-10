package main

import (
	"fmt"
)

func changePointValue(p *int) {
	fmt.Println("point value before change:", *p)
	*p += 10
}

func changeSliceValue(s *[]int) {
	for i := range *s {
		(*s)[i] *= 2
	}
}
