package main

import "fmt"

func main() {
	a, b := 0, 0
	fmt.Scan(&a, &b)

	for i := a; i <= b; i++ {
		if 100 % i == 0 {
			fmt.Println("Yes")
			return
		}
	}

	fmt.Println("No")
}
