package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)

	var N, K int
	sc.Scan()
	fmt.Sscanf(sc.Text(), "%d %d", &N, &K)

	s := make([]int, N+1)
	for i := 1; i <= N; i++ {
		sc.Scan()
		fmt.Sscanf(sc.Text(), "%d", &s[i])
	}

	for i := 1; i <= N; i++ {
		if s[i] == 0 {
			fmt.Println(N)
			return
		}
	}

	/*
		k = 6
		[4,3,1,1,2,10,2]
		========================
		right:_1_left:_1_sum:_4
		right:_1_left:_1_sum:_4
		-========================
		right:_2_left:_1_sum:_12
		right:_2_left:_2_sum:_3
		-========================
		right:_3_left:_2_sum:_3
		right:_3_left:_2_sum:_3
		-========================
		right:_4_left:_2_sum:_3
		right:_4_left:_2_sum:_3
		-========================
		right:_5_left:_2_sum:_6
		right:_5_left:_2_sum:_6
		-========================
		right:_6_left:_2_sum:_60
		right:_6_left:_7_sum:_1
		-========================
		right:_7_left:_7_sum:_2
		right:_7_left:_7_sum:_2
	*/

	ans := 0
	sum := 1
	left := 1
	for right := 1; right <= N; right++ {
		sum *= s[right]

		for left <= right && sum > K {
			sum /= s[left]
			left++
		}
		ans = max(ans, right-left+1)
	}

	fmt.Println(ans)
}

func max(a, b int) int {
	if a < b {
		return b
	}

	return a
}
