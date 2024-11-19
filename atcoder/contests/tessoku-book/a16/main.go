package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Buffer(make([]byte, 1000), 400000)

	var N int
	sc.Scan()
	fmt.Sscanf(sc.Text(), "%d", &N)

	// 1 indexed なので +1
	A := make([]int, N+1)
	sc.Scan()
	line := strings.Fields(sc.Text())
	// 2 <= i <= N なので +2
	for i := 2; i <= N; i++ {
		A[i], _ = strconv.Atoi(line[i-2])
	}

	// 1 indexed なので +1
	B := make([]int, N+1)
	sc.Scan()
	line = strings.Fields(sc.Text())
	// 3 <= i <= N なので +2
	for i := 3; i <= N; i++ {
		B[i], _ = strconv.Atoi(line[i-3])
	}

	dp := make([]int, N+1)
	dp[1] = 0    // 最初の部屋までは 0 分で移動可能
	dp[2] = A[2] // 2 番目の部屋へは部屋1からの道のみ
	for i := 3; i <= N; i++ {
		dp[i] = min(dp[i-1]+A[i], dp[i-2]+B[i])
	}

	fmt.Println(dp[N])
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
