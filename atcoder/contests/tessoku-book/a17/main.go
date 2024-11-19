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

	A := make([]int, N+1)
	sc.Scan()
	line := strings.Fields(sc.Text())
	for i := 2; i <= N; i++ {
		A[i], _ = strconv.Atoi(line[i-2])
	}

	B := make([]int, N+1)
	sc.Scan()
	line = strings.Fields(sc.Text())
	for i := 3; i <= N; i++ {
		B[i], _ = strconv.Atoi(line[i-3])
	}

	dp := make([]int, N+1)
	dp[1] = 0
	dp[2] = A[2]
	for i := 3; i <= N; i++ {
		dp[i] = min(A[i]+dp[i-1], B[i]+dp[i-2])
	}

	answer := []int{}
	place := N
	for {
		answer = append(answer, place)
		if place == 1 {
			break
		}

		if dp[place-1]+A[place] == dp[place] {
			place -= 1
		} else {
			place -= 2
		}
	}

	answer = reverse(answer)
	fmt.Println(len(answer))

	strs := make([]string, len(answer))
	for i, v := range answer {
		strs[i] = strconv.Itoa(v)
	}
	fmt.Println(strings.Join(strs, " "))
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func reverse(s []int) []int {
	result := make([]int, len(s))
	for i := range s {
		result[i] = s[len(s)-1-i]
	}

	return result
}
