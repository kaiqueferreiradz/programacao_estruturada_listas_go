package main

import (
	"fmt"
)

func main() {

	var a int
	fmt.Println("N:")
	fmt.Scanln(&a)

	b := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
	c := []int{}

	for i := 0; i < a; i++ {
		aa := 0
		fmt.Scanln(&aa)
		c = append(c, aa)
	}

	//fmt.Println(b)
	x := troca(b, len(b), c, len(c))
	fmt.Println(x)
}

func troca(a []string, b int, c []int, d int) string {

	ret := ""

	for i := 0; i < d; i++ {
		ret = ret + a[c[i]]
	}

	return ret
}
