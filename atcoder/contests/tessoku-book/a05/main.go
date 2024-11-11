package main

import "fmt"

func main() {
	var (
		n int
		k int
	)
	fmt.Scan(&n, &k)

	result := 0

	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			target := k - i - j
			if target >= 1 && target <= n {
				result++
			}
		}
	}

	fmt.Println(result)
}
