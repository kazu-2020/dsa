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
	sc.Scan()
	inputs := strings.Fields(sc.Text())

	result := 0
	for _, n := range inputs {
		num, _ := strconv.Atoi(n)
		result += num
	}

	fmt.Printf("%d\n", result)
}
