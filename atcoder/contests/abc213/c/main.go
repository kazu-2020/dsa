package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func compress(s []int) []int {
	clone := make([]int, len(s))
	copy(clone, s)

	sort.Slice(clone, func(i, j int) bool {
		return clone[i] <= clone[j]
	})
	clone = uniq(clone)

	res := make([]int, len(s))
	for i := 0; i < len(s); i++ {
		res[i] = lowerBound(clone, s[i])
	}

	return res
}

func uniq(s []int) []int {
	m := map[int]bool{}
	res := []int{}

	for i := 0; i < len(s); i++ {
		if !m[s[i]] {
			m[s[i]] = true
			res = append(res, s[i])
		}
	}

	return res
}

func lowerBound(s []int, t int) int {
	result := sort.Search(len(s), func(i int) bool {
		return s[i] >= t
	})

	return result
}

func main() {
	sc := bufio.NewScanner(os.Stdin)

	var H, W, N int
	sc.Scan()
	fmt.Sscanf(sc.Text(), "%d %d %d", &H, &W, &N)

	A := make([]int, N)
	B := make([]int, N)

	for i := 0; i < N; i++ {
		sc.Scan()
		fmt.Sscanf(sc.Text(), "%d %d", &A[i], &B[i])
	}

	ca := compress(A)
	cb := compress(B)
	for i := 0; i < N; i++ {
		fmt.Printf("%d %d\n", ca[i]+1, cb[i]+1)
	}
}
