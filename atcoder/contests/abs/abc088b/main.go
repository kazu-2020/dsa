package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var N int
	fmt.Scan(&N)
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	line := strings.Fields(sc.Text())
	a := make([]int, N)
	for i := range line {
		a[i], _ = strconv.Atoi(line[i])
	}
	// 1.20.6 では使えない
	slices.SortFunc(a, func(a, b int) int {
		return b - a
	})
	sort.Sort(sort.Reverse(sort.IntSlice(a)))

	alice, bob := 0, 0
	for i := 0; i < N; i++ {
		if i % 2 == 0 {
			alice += a[i]
		} else {
			bob += a[i]
		}
	}

	fmt.Println(alice - bob)
}
