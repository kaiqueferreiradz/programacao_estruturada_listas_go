package main

import (
	"fmt"
)

func main() {

	b := ""
	c := ""

	fmt.Println("N 1 -:")
	fmt.Scanln(&b)
	fmt.Println("N 2 -:")
	fmt.Scanln(&c)

	fmt.Println(len(b) == len(c))

	a := []int{}

	for i := 0; i < len(b); i++ {
		a = append(a, int(b[i]))
	}

	cc := 0
	for i := 0; i < len(b); i++ {
		for j := 0; j < len(b); j++ {
			if int(c[i]) == a[j] {
				a[j] = -1
				cc++
				break
			}
		}
	}

	fmt.Println(cc == len(c))

}
