package main

import "fmt"

func main() {

	var a int
	fmt.Println("N:")
	fmt.Scanln(&a)
	fmt.Println(test(a))
}

func test(a int) int {

	for i := 0; i <= a; i++ {
		for j := 0; j <= a; j++ {
			if i*j == a {
				fmt.Println(i, j)
			}
		}
	}
	for i := 0; i >= -a; i-- {
		for j := 0; j >= -a; j-- {
			if i*j == a {
				fmt.Println(i, j)
			}
		}
	}
	for i := 0; i <= -a; i++ {
		for j := 0; j >= a; j-- {
			if i*j == a {
				fmt.Println(i, j)
			}
		}
	}
	for i := 0; i >= a; i-- {
		for j := 0; j <= -a; j++ {
			if i*j == a {
				fmt.Println(i, j)
			}
		}
	}
	return 0
}
