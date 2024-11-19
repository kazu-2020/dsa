package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)

	var N int
	sc.Scan()
	fmt.Sscanf(sc.Text(), "%d", &N)

	X := make([]int, N)
	sc.Scan()
	line := strings.Fields(sc.Text())
	for i := range line {
		X[i], _ = strconv.Atoi(line[i])
	}

	min_num, max_num := minmax(X)

	ans := math.MaxInt64
	for p := min_num; p <= max_num; p++ {
		sum := 0
		for i := range X {
			sum += ((X[i] - p) * (X[i] - p))
		}

		if ans > sum {
			ans = sum
		}
	}

	fmt.Println(ans)
}

func minmax(nums []int) (min int, max int) {
	min = nums[0]
	max = nums[0]

	for _, num := range nums {
		if num < min {
			min = num
		}
		if num > max {
			max = num
		}
	}

	return min, max
}
