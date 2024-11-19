package main

import "fmt"

func main() {
	var M1, D1, M2, D2 int
	fmt.Scan(&M1, &D1)
	fmt.Scan(&M2, &D2)

	if M1 == M2 {
		fmt.Println(0)
	} else {
		fmt.Println(1)
	}
}
