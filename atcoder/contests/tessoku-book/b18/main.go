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

	for i := 1; i <= N; i++ {
		for j := 0; j <= S; j++ {
			if j < A[i] {
				if dp[i-1][j] {
					dp[i][j] = true
				}
			} else {
				if dp[i-1][j] || dp[i-1][j-A[i]] {
					dp[i][j] = true
				}
			}
		}
	}

	if !dp[N][S] {
		fmt.Println(-1)
		return
	}

	ans := []int{}
	sum := S
	for i := N; i > 0; i-- {
		if sum == 0 {
			break
		}

		if dp[i-1][sum] {
			sum -= 0
		} else {
			sum -= A[i]
			ans = append(ans, i)
		}
	}

	ans = reverse(ans)
	strs := make([]string, len(ans))
	for i, v := range ans {
		strs[i] = strconv.Itoa(v)
	}

	fmt.Println(len(strs))
	fmt.Println(strings.Join(strs, " "))
}

func reverse(nums []int) []int {
	result := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		result[i] = nums[len(nums)-1-i]
	}

	return result
}
