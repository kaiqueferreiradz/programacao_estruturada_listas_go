package main

import (
	"fmt"
)

func main() {

	fmt.Println("N:")

	b := []int{1, 2, 5, 8}
	c := []int{1, 2, 3, 4, 5, 6, 7}

	//fmt.Println(b)
	x := troca(b, len(b), c, len(c))
	fmt.Println(x)
}

func troca(a []int, b int, c []int, d int) int {

	cc := []int{}
	cc = c
	f := 0
	for k := 0; k < b; k++ {
		for i := 0; i < d; i++ {
			if a[k] == c[i] {
				f = 1
			}
		}
		if f == 0 {
			cc = append(cc, a[k])
		} else {
			f = 0
		}
	}
	fmt.Println(cc)
	return len(cc)
}
