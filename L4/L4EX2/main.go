package main

import (
	"fmt"
)

func main() {

	fmt.Println("N:")

	b := [5]int{}
	c := [5]int{}

	for i := 0; i < 5; i++ {
		fmt.Scanln(&b[i])
	}

	for i := 0; i < 5; i++ {
		c[i] = b[4-i]
	}
	fmt.Println(c)

}
