package main

import "fmt"

func main() {

	for i := 1; i <= 20; i++ {
		fmt.Printf("%d! = %d \n", i, fat(i))
	}
}

func fat(a int) int {

	ret := 1
	for i := 1; i <= a; i++ {
		ret = ret * i
	}
	return ret
}
