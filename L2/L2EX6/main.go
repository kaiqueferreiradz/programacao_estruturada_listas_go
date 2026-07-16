package main

import "fmt"

func main() {

	fmt.Println("N:")
	var n int
	fmt.Scanln(&n)

	var a int
	fmt.Println("N", 1, ":")
	fmt.Scanln(&a)
	cc := a
	ret := true
	for i := 2; i <= n; i++ {
		fmt.Println("N", i, ":")
		var a int
		fmt.Scanln(&a)
		if a < cc {
			ret = false
		}
		cc = a
	}

	if ret {
		fmt.Println("Ordenado")
	} else {
		fmt.Println("Nao Ordenado")
	}

}
