package main

import (
	"fmt"
)

func main() {

	fmt.Println("N:")

	b := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(b)
	troca(b, 0, 3)
	fmt.Println(b)
}

func troca(a []int, i int, j int) {
	aux := a[i]
	a[i] = a[j]
	a[j] = aux
}
