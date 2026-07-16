package main

import (
	"fmt"
)

func main() {

	fmt.Println("N:")

	b := []int{1, 2, 5, 4}
	c := []int{3, 1, 4, 2}

	//fmt.Println(b)
	x := troca(b, len(b), c, len(c))
	fmt.Println(x)
}

func troca(a []int, b int, c []int, d int) int {

	cc := 1
	for k := 0; k < b; k++ {
		ff := 0
		for i := 0; i < d; i++ {
			if a[k] == c[i] {
				ff = 1
			}
		}
		if ff == 1 {
			ff = 0
		} else {
			cc = 0
		}
	}
	return cc
}
