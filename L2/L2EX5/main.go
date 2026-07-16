package main

import "fmt"

func main() {

	var n int
	fmt.Println("N:")
	fmt.Scanln(&n)

	for j := 1; j <= n; j++ {
		c := 0
		for i := 1; i <= j; i++ {
			c = c + i
		}
		fmt.Println(c)
	}
}
