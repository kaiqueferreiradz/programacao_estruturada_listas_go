package main

import "fmt"

func main() {
	fmt.Println("Entrada X Y:")
	var x, y int

	fmt.Scanln(&x, &y)

	if x < 0 && y > 0 {
		fmt.Println("Primeiro")
	} else if x > 0 && y > 0 {
		fmt.Println("Segundo")
	} else if x > 0 && y < 0 {
		fmt.Println("Terceiro")
	} else if x < 0 && y < 0 {
		fmt.Println("Quarto")
	} else if x == 0 || y == 0 {
		fmt.Println("NAN")
	}
}
