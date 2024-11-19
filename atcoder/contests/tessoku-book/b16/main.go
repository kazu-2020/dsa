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
	sc.Buffer(make([]byte, 1000), 1000000)

	var N int
	sc.Scan()
	fmt.Sscanf(sc.Text(), "%d", &N)

	h_list := make([]int, N+1)
	sc.Scan()
	line := strings.Fields(sc.Text())
	for i := 1; i <= N; i++ {
		h_list[i], _ = strconv.Atoi(line[i-1])
	}

	dp := make([]int, N+1)
	dp[1] = 0
	dp[2] = abs(h_list[2] - h_list[1])

	for i := 3; i <= N; i++ {
		dp[i] = min(abs(h_list[i-1]-h_list[i])+dp[i-1], abs(h_list[i-2]-h_list[i])+dp[i-2])
	}

	fmt.Println(dp[N])
}

func abs(n int) int {
	if n < 0 {
		return -n
	}

	return n
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
