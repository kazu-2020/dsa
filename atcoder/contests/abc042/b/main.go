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
	// Scan() を呼ぶと、1 行読み込む
	sc.Scan()
	nums := splitToInts(sc.Text())
	n := nums[0]
	stringsList := make([]string, n)
	for i := 0; i < n; i++ {
		sc.Scan()
		stringsList[i] = sc.Text()
	}

	sort.Strings(stringsList)
	fmt.Println(strings.Join(stringsList, ""))
}

func splitToInts(s string) []int {
	fields := strings.Fields(s)
	nums := make([]int, len(fields))

	for i, f := range fields {
		nums[i], _ = strconv.Atoi(f)
	}

	return nums
}
