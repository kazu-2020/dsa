package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Buffer(make([]byte, 1000000), 10000000)

	var N, M int
	sc.Scan()
	fmt.Sscanf(sc.Text(), "%d %d", &N, &M)

	A, B := make([]int, N), make([]int, M)
	sc.Scan()
	line := strings.Fields(sc.Text())
	for i := 0; i < N; i++ {
		A[i], _ = strconv.Atoi(line[i])
	}
	sc.Scan()
	line = strings.Fields(sc.Text())
	for i := 0; i < M; i++ {
		B[i], _ = strconv.Atoi(line[i])
	}

	sort.Slice(A, func(a, b int) bool {
		return A[a] < A[b]
	})
	sort.Slice(B, func(a, b int) bool {
		return B[a] < B[b]
	})

	ai, bi := 0, 0
	ans := math.MaxInt64
	for {
		if N <= ai || M <= bi {
			break
		}
		ans = min(ans, abs(A[ai]-B[bi]))
		if A[ai] < B[bi] {
			ai++
		} else {
			bi++
		}
	}

	fmt.Println(ans)
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func abs(n int) int {
	if n < 0 {
		return -n
	}

	return n
}
