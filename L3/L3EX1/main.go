package main

import "fmt"

func main() {

	for i := 0; i <= 10; i++ {
		w := pot(2, i)
		fmt.Println(w)
	}

}

func pot(a int, b int) int {

	ret := 1
	for i := 0; i < b; i++ {
		ret = ret * a
	}

	return ret
}
