package main

import (
	"fmt"
	"math"
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
		cc = cc + a[k]
	}

	dd := 0
	for k := 0; k < b; k++ {
		dd = dd + int(math.Pow(float64(a[k]-cc/b), 2.0))/(b-1)
	}

	return dd
}
