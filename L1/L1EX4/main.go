package main

import "fmt"

func main() {

	var a int
	var b int
	fmt.Println("N1:")
	fmt.Scanln(&a)
	fmt.Println("N2")
	fmt.Scanln(&b)

	a = a + b
	b = a - b
	a = a - b

	fmt.Printf("N1: %d \n", a)
	fmt.Printf("N2: %d \n", b)
}
