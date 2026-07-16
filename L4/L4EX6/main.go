package main

import (
	"fmt"
)

func main() {

	fmt.Println("N:")

	b := []int{1, 2, 5, 6, 7}
	c := []int{1, 2, 3, 4, 5, 6, 7}

	fmt.Println(b)
	x := troca(b, len(b), c, len(c))
	fmt.Println(x)
}

func troca(a []int, b int, c []int, d int) int {

	cc := 1
	for k := 0; k < b; k++ {
		f := 0
		for i := 0; i < d; i++ {
			if a[k] == c[i] {
				f = 1
			}
		}
		if f == 0 {
			cc = 0
		}
	}
	return cc
}
