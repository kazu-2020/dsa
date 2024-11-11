package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var a int64
	fmt.Scan(&a)

	binary := strconv.FormatInt(a, 2)
	prefix := strings.Repeat("0", 10 - len(binary))
	fmt.Println(prefix + binary)
}
