package main

import (
	"fmt"
)

func main() {

	fmt.Println("N:")

	b := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(b)
	troca(b, len(b))
	fmt.Println(b)
}

func troca(a []int, b int) {

	for k := 0; k < (b / 2); k++ {
		aux := a[k]
		a[k] = a[b-k-1]
		a[b-k-1] = aux
	}
}
