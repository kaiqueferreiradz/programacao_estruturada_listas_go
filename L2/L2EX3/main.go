package main

import "fmt"

func main() {

	fmt.Println("N:")
	var n int
	fmt.Scanln(&n)

	c := 0
	for i := 0; i <= n; i++ {
		c = c + i
	}
	fmt.Println(c)

}
