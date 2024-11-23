package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var N int
	fmt.Scan(&N)
	sc := bufio.NewScanner(os.Stdin)
	sc.Buffer(make([]byte, 1000000), 1000000)

	A := make([]int, N)
	sc.Scan()
	line := strings.Fields(sc.Text())
	for i := range line {
		A[i], _ = strconv.Atoi(line[i])
	}

}
