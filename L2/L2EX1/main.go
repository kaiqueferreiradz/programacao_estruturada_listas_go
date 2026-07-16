package main

import "fmt"

func main() {
	a := 1
	b := 0

	fmt.Println("1. Sopa")
	fmt.Println("2. Peixe")
	fmt.Println("3. Pasta")
	fmt.Println("4. Arroz")
	fmt.Println("5. sair")

	for a == 1 {

		fmt.Scanln(&b)

		if b == 1 {
			fmt.Println("Sopa")
		} else if b == 2 {
			fmt.Println("Peixe")
		} else if b == 3 {
			fmt.Println("Pasta")
		} else if b == 4 {
			fmt.Println("Arroz")
		} else if b == 5 {
			fmt.Println("Sair")
			a = 0
		} else {
			fmt.Println("Invalido")
		}
	}
}
