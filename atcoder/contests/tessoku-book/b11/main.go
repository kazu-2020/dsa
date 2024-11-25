package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Buffer(make([]byte, 1000000), 10000000)

	var N int
	sc.Scan()
	fmt.Sscanf(sc.Text(), "%d", &N)

	A := make([]int, N)
	sc.Scan()
	line := strings.Fields(sc.Text())
	for i := 0; i < N; i++ {
		A[i], _ = strconv.Atoi(line[i])
	}

	// 昇順ソート
	sort.Slice(A, func(i, j int) bool {
		return A[i] < A[j]
	})

	var Q int
	sc.Scan()
	fmt.Sscanf(sc.Text(), "%d", &Q)

	for i := 1; i <= Q; i++ {
		var X int
		sc.Scan()
		fmt.Sscanf(sc.Text(), "%d", &X)

		fmt.Println(lowerBound(A, X))
	}
}

func lowerBound(s []int, target int) int {
	result := sort.Search(len(s), func(i int) bool { return s[i] >= target })

	return result
}
