package main

import "fmt"

func main() {

	fmt.Println("Ano:")
	var a int
	fmt.Scanln(&a)

	if a%400 == 0 {
		fmt.Println("Bissexto")
	} else if a%400 != 0 && a%100 != 0 && a%4 == 0 {
		fmt.Println("Bissexto")
	} else {
		fmt.Println("Nao Bissexto")
	}
}
