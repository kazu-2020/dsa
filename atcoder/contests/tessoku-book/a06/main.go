package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var (
		n int
		q int
	)
	fmt.Scan(&n, &q)
	sc := bufio.NewScanner(os.Stdin)
	// 1 <= n <= 10^5 のため、バッファサイズを大きくしておく
	sc.Buffer(make([]byte, 1000000), 1000000)
	// 日数読み込み
	sc.Scan()
	line := strings.Fields(sc.Text())
	A := make([]int, n + 1)
	A[0] = 0
	for i := 1; i <= n; i++ {
		A[i], _ = strconv.Atoi(line[i - 1])
	}
	// 累積和
	S := make([]int, n + 1)
	S[0] = 0
	for i := 1; i <= n; i++ {
		S[i] = S[i - 1] + A[i]
	}

	for i := 1; i <= q; i++ {
		sc.Scan()
		line := strings.Fields(sc.Text())
		L, _ := strconv.Atoi(line[0])
		R, _ := strconv.Atoi(line[1])

		fmt.Println(S[R] - S[L - 1])
	}
}
