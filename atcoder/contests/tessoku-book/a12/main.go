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
	sc.Buffer(make([]byte, 10000000), 10000000)

	var N, K int
	sc.Scan()
	fmt.Sscanf(sc.Text(), "%d %d", &N, &K)

	A := make([]int, N)
	sc.Scan()
	line := strings.Fields(sc.Text())
	for i := 0; i < N; i++ {
		A[i], _ = strconv.Atoi(line[i])
	}

	left, right := 0, 1000000000
	for left < right {
		mid := (left + right) / 2
		if check(A, K, mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}

	fmt.Println(left)
}

func check(A []int, K int, t int) bool {
	sum := 0
	for _, v := range A {
		sum += t / v
	}

	return sum >= K
}
