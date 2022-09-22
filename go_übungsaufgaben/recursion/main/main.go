package main

import "fmt"

func main() {
	fmt.Println(fac(5))
}

func fac(x int) int {
	if x == 0 {
		return 1
	}

	return fac(x-1) * x
}
