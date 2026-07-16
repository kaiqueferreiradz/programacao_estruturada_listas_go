package main

import (
	"fmt"
)

func main() {

	var a int
	fmt.Println("N:")
	fmt.Scanln(&a)

	fmt.Println(test2(a, a))
}

func test(a int, b int) int {

	fmt.Printf(". ")
	for i := 0; i < b-a; i++ {
		fmt.Printf("-")
	}
	fmt.Printf("\n")
	if a > 0 {
		n := test(a-1, b)
		if b-n > 1 {
			fmt.Printf(". ")
		}
		for j := 1; j < b-n; j++ {
			fmt.Printf("-")
		}
		if b-n > 1 {
			fmt.Printf("\n")
		}

	}
	return a

}

func test2(a int, b int) int {

	test(b-a, b-a)
	if a > -0 {
		n := test2(a-1, b)
		test(b-n, b-n)
	}
	return a

}
