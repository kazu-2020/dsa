package main

import "fmt"

func main() {
	var D, N int
	fmt.Scan(&D)
	fmt.Scan(&N)

	L := make([]int, N + 1)
	R := make([]int, N + 1)

	L[0] = 0
	R[0] = 0
	for i := 1; i <= N; i++ {
		fmt.Scan(&L[i], &R[i])
	}

	// 前日比を計算
	S := make([]int, D + 2)
	S[0] = 0
	for i := 1; i <= N; i++ {
		S[L[i]] += 1
		S[R[i] + 1] -= 1
	}

	// 累積和を計算
	answer := make([]int, D + 1)
	answer[0] = 0
	for i := 1; i <= D; i++ {
		answer[i] = answer[i - 1] + S[i]
		fmt.Println(answer[i])
	}
}
