package main

import (
	"fmt"
	"strconv"
)

func main() {
	var input string
	fmt.Scan(&input)

	n, _ := strconv.Atoi(input)
	fmt.Println(n * n)
}
