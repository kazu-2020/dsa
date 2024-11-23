package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	var a, b string
	fmt.Scan(&a, &b)

	t, _ := strconv.Atoi(a + b)
	result := math.Sqrt(float64(t))

	if int(result)*int(result) == t {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
