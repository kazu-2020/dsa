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
	var N, K int

	sc.Scan()
	fmt.Sscanf(sc.Text(), "%d", &N)
	sc.Scan()
	fmt.Sscanf(sc.Text(), "%d", &K)

	X := make([]int, N)
	sc.Scan()
	line := strings.Fields(sc.Text())
	for i := 0; i < N; i++ {
		X[i], _ = strconv.Atoi(line[i])
	}

	sum := 0
	for i := 0; i < N; i++ {
		curr := X[i]
		sum += min(curr, abs(K-curr)) * 2
	}

	fmt.Println(sum)
}

func abs(num int) int {
	if num < 0 {
		return -num
	}

	return num
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
