package main

import (
	"fmt"
)

func main() {
	var N, A, B int
	fmt.Scan(&N, &A, &B)

	ans := 0
	for i := 1; i <= N; i++ {
		result := sum(i)

		if A <= result && result <= B {
			ans += i
		}
	}

	fmt.Println(ans)
}

func sum(num int) int {
	sum := 0
	for num > 0 {
		sum += num % 10
		num /= 10
	}

	return sum
}
