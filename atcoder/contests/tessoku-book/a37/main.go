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
	sc.Buffer(make([]byte, 1000000000), 1000000000)

	var N, M, B int
	sc.Scan()
	fmt.Sscanf(sc.Text(), "%d %d %d", &N, &M, &B)

	A, C := make([]int, N), make([]int, M)
	sc.Scan()
	for i, v := range strings.Fields(sc.Text()) {
		A[i], _ = strconv.Atoi(v)
	}
	sc.Scan()
	for i, v := range strings.Fields(sc.Text()) {
		C[i], _ = strconv.Atoi(v)
	}

	ans := 0
	for _, v := range A {
		ans += M * v
	}
	ans += B * N * M
	for _, v := range C {
		ans += N * v
	}

	fmt.Println(ans)
}
