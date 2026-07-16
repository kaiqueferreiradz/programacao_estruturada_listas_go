package main

import (
	"fmt"
)

func main() {

	b := ""
	c := ""

	fmt.Println("TXT -:")
	fmt.Scanln(&b)
	fmt.Println("P -:")
	fmt.Scanln(&c)

	a := []int{}
	cc := 1
	for i := 0; i < len(b); i++ {
		for j := 0; j < len(c); j++ {
			if j+i < len(b) {
				fmt.Println(c[j], "xx", b[j+i])
			}
			if j+i < len(b) && c[j] != b[j+i] {
				cc = 0
			}
		}
		if cc == 1 {
			a = append(a, i)
		}
		cc = 1
	}

	fmt.Println(a)

}
