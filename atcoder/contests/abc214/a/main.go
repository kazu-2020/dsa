package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)

	if 1 <= N && N <= 125 {
		fmt.Println(4)
		return
	} else if 126 <= N && N <= 211 {
		fmt.Println(6)
		return
	}

	fmt.Println(8)
}
