package main

import "fmt"

func main() {

	var a int
	fmt.Println("N:")
	fmt.Scanln(&a)
	fmt.Println(fib(a))

}

func fib(a int) int {
	ret := 0
	i := 1
	for i < a {
		aux := i
		i = i + ret
		ret = aux
	}
	return ret
}
