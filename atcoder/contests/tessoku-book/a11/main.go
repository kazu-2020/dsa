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
	sc.Buffer(make([]byte, 100000), 10000000)

	var N, X int
	sc.Scan()
	fmt.Sscanf(sc.Text(), "%d %d", &N, &X)

	A := make([]int, N+1) // 1 indexed
	sc.Scan()
	line := strings.Fields(sc.Text())
	for i := 1; i <= N; i++ {
		A[i], _ = strconv.Atoi(line[i-1])
	}

	search := func(nums []int, target int) int {
		l := 1
		r := N

		for l <= r {
			M := (l + r) / 2
			if target < nums[M] {
				r = M - 1
			}
			if target == nums[M] {
				return M
			}
			if target > nums[M] {
				l = M + 1
			}
		}

		return -1
	}

	ans := search(A, X)
	fmt.Println(ans)
}
