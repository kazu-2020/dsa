package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)

	var N, W int
	sc.Scan()
	fmt.Sscanf(sc.Text(), "%d %d", &N, &W)

	w := make([]int, N+1)
	v := make([]int, N+1)
	for i := 1; i <= N; i++ {
		sc.Scan()
		fmt.Sscanf(sc.Text(), "%d %d", &w[i], &v[i])
	}

	// 縦: 選択した品物の数, 横: 総重量
	dp := make([][]int, N+1)
	for i := 0; i <= N; i++ {
		dp[i] = make([]int, W+1)
	}

	dp[0][0] = 0
	for i := 1; i <= W; i++ {
		dp[0][i] = math.MinInt64
	}
	for i := 1; i <= N; i++ {
		for j := 0; j <= W; j++ {
			if j < w[i] {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i-1][j-w[i]]+v[i])
			}
		}
	}

	ans := 0
	for i := 0; i <= W; i++ {
		ans = max(ans, dp[N][i])
	}

	fmt.Println(ans)
}

func max(a, b int) int {
	if a < b {
		return b
	}

	return a
}
