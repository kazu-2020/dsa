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

	var N, M, C int
	sc.Scan()
	fmt.Sscanf(sc.Text(), "%d %d %d", &N, &M, &C)

	B := make([]int, M)
	sc.Scan()
	line := strings.Fields(sc.Text())
	for i := range line {
		B[i], _ = strconv.Atoi(line[i])
	}

	ans := 0
	for i := 1; i <= N; i++ {
		sum := 0
		sc.Scan()
		line := strings.Fields(sc.Text())

		for j := range line {
			A, _ := strconv.Atoi(line[j])
			sum += A * B[j]
		}

		if sum+C > 0 {
			ans++
		}
	}

	fmt.Println(ans)
}
