package main

import "fmt"

func main() {

	fmt.Println("3 numeros:")
	var a int
	var b int
	var c int
	fmt.Scanln(&a, &b, &c)

	if a > b {
		aux := b
		b = a
		a = aux
	}

	if a > c {
		aux := c
		c = a
		a = aux
	}

	if b > c {
		aux := c
		c = b
		b = aux
	}

	fmt.Printf(" A: %d- B: %d - C: %d ", a, b, c)

}
