package main

import (
	"fmt"
	"strconv"
)

func main() {
	var binary string
	fmt.Scan(&binary)

	num, _ := strconv.ParseInt(binary, 2, 64)
	fmt.Println(num)
}
