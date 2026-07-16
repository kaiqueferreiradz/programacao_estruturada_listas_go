package main

import "fmt"

func main() {

	fmt.Println("F ou C:")

	var a string
	var res float64

	fmt.Scanln(&a)

	if a == "F" || a == "f" {
		fmt.Println("F:")
		var n float64
		fmt.Scanln(&n)

		res = (5 * (n - 32)) / 9
		fmt.Println("C: ", res)

	}
	if a == "C" || a == "c" {
		fmt.Println("C:")
		var n float64
		fmt.Scanln(&n)

		res = n*1.8 + 32
		fmt.Println("F: ", res)

	}
}
