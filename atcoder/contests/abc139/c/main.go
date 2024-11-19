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
	sc.Buffer(make([]byte, 1000), 10000000)
	sc.Scan()
	var N int
	fmt.Sscanf(sc.Text(), "%d", &N)

	// マスの数が1の場合はジャンプする必要がない
	if N == 1 {
		fmt.Println(0)
		return
	}

	H := make([]int, N)
	sc.Scan()
	line := strings.Fields(sc.Text())
	for i := 0; i < N; i++ {
		H[i], _ = strconv.Atoi(line[i])
	}

	curr  := 0 // 現在どれだけジャンプしているか
	ans   := 0 // 最も長いジャンプの回数
	for i := 1; i < N; i++ {
		if H[i] <= H[i-1] {
			curr++
			ans = max(ans, curr)
		} else {
			curr = 0
		}
	}

	fmt.Println(ans)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
