package main

import (
	"fmt"
	"strings"
)

const TARGET = "1 1 1 2 1 2 1 5 2 2 1 5 1 2 1 14 1 5 1 5 2 2 1 15 2 2 5 4 1 4 1 51"

func main() {
	var K int
	fmt.Scan(&K)

	ans := strings.Fields(TARGET)[K-1]
	fmt.Println(ans)
}
