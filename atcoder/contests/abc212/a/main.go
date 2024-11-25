package main

import "fmt"

func main() {
	var A, B int
	fmt.Scan(&A, &B)

	if 0 < A && B == 0 {
		fmt.Println("Gold")
		return
	}

	if A == 0 && 0 < B {
		fmt.Println("Silver")
		return
	}

	fmt.Println("Alloy")
}
