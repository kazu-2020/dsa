package main

import "fmt"

func main() {
	var A, B int
	fmt.Scan(&A, &B)

	n := 0
	for (A-1)*n+1 < B {
		n++
	}

	fmt.Println(n)
}
