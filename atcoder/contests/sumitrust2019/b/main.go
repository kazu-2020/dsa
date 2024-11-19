package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)

	for i := 1; i <= N; i++ {
		curr := int(float64(i) * 1.08)

		if curr == N {
			fmt.Println(i)
			return
		}
	}

	fmt.Println(":(")
}
