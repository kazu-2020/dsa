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

	A := make([]Score, N)
	sc.Scan()
	line := strings.Fields(sc.Text())
	for i := 0; i < N; i++ {
		val, _ := strconv.Atoi(line[i])
		A[i] = Score{i + 1, val}
	}

	sort.Slice(A, func(i, j int) bool {
		return A[i].value > A[j].value
	})

	fmt.Println(A[1].pos)
}

type Score struct {
	pos   int
	value int
}
