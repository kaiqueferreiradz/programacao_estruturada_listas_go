package main

import "fmt"

func main() {

	fmt.Println("N:")
	var n int
	fmt.Scanln(&n)

	for i := 2; i <= n; i++ {
		if n%i == 0 {
			fmt.Println(i)
		}
	}

}
