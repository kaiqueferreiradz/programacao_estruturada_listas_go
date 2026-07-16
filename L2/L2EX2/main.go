package main

import "fmt"

func main() {

	var a, c int
	fmt.Println("Entrar A^B:")
	fmt.Scanln(&a, &c)

	b := 1
	for i := 0; i < c; i++ {
		b = b * a
	}

	fmt.Println(b)

}
