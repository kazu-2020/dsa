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
	sc.Buffer(make([]byte, 10000), 10000)

	var N, S int
	sc.Scan()
	fmt.Sscanf(sc.Text(), "%d %d", &N, &S)

	A := make([]int, N+1)
	sc.Scan()
	line := strings.Fields(sc.Text())
	for i := 1; i <= N; i++ {
		A[i], _ = strconv.Atoi(line[i-1])
	}

	dp := make([][]bool, N+1)
	for i := 0; i <= N; i++ {
		dp[i] = make([]bool, S+1)
	}
	dp[0][0] = true
	for i := 1; i <= S; i++ {
		dp[0][i] = false
	}

	for i := 1; i <= N; i++ {
		for j := 0; j <= S; j++ {
			if j < A[i] {
				if dp[i-1][j] {
					dp[i][j] = true
				} else {
					dp[i][j] = false
				}
			}
			if j >= A[i] {
				if dp[i-1][j] || dp[i-1][j-A[i]] {
					dp[i][j] = true
				} else {
					dp[i][j] = false
				}
			}
		}
	}

	if dp[N][S] {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
