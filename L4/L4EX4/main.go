package main

import (
	"fmt"
)

func main() {

	fmt.Println("N:")

	b := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(b)
	x := troca(b, len(b))
	fmt.Println(x)
}

func troca(a []int, b int) int {

	cc := 0
	for k := 0; k < b; k++ {
		if k%2 == 0 {
			cc = cc + a[k]
		}
	}
	return cc
}
