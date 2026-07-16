package main

import "fmt"

func main() {
	var a int
	fmt.Println("N:")
	fmt.Scanln(&a)
	fmt.Println(pri(a))
}

func pri(a int) int {

	ret := 0
	for i := a; i > 0; i-- {

		v := true
		for j := 2; j < i; j++ {
			if i%j == 0 {
				v = false
			}
		}
		if v {
			ret = i
			i = 0
		}
	}
	return ret
}
