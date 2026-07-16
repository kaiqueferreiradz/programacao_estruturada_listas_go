package main

import (
	"fmt"
)

func main() {

	b := ""

	fmt.Println("TXT -:")
	fmt.Scanln(&b)

	cc := 1
	for i := 0; i < len(b); i++ {
		if b[i] != b[len(b)-1-i] {
			cc = 0
		}
	}

	if cc == 1 {
		fmt.Println("palin")
	} else {
		fmt.Println("n-palin")
	}

}
