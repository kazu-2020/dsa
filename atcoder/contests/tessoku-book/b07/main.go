package main

import "fmt"

func main() {
	var T, N int
	fmt.Scan(&T, &N)

	L := make([]int, N+1)
	R := make([]int, N+1)
	L[0] = 0
	R[0] = 0
	for i := 1; i <= N; i++ {
		fmt.Scan(&L[i], &R[i])
	}

	// 前の時刻より従業員がどれだけ増えたかを記録する
	S := make([]int, T+2)
	for i := 1; i <= N; i++ {
		S[L[i]] += 1
		S[R[i]] -= 1
	}

	// 累積和を取る
	answer := make([]int, T+1)
	answer[0] = S[0]
	for i := 1; i <= T; i++ {
		answer[i] = answer[i-1] + S[i]
	}

	for i := 0; i < T; i++ {
		fmt.Println(answer[i])
	}
}
