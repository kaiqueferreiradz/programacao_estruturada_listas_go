package main

import "fmt"

func main() {

	var a, b, c int

	fmt.Println("Lados:")
	fmt.Scanln(&a, &b, &c)

	if a == b && b == c {
		fmt.Println("Equilatero e Isosceles")
	} else if (a == b && b != c) || (a == c && c != b) || (b == c && c != a) {
		fmt.Println("Isosceles")
	} else {
		fmt.Println("Escaleno")
	}

}
