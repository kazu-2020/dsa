package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var H, W int
	fmt.Scan(&H, &W)

	sc := bufio.NewScanner(os.Stdin)
	sc.Buffer(make([]byte, 10000), 1000000)

	X := make([][]int, H+1)
	for i := 1; i <= H; i++ {
		sc.Scan()
		line := strings.Fields(sc.Text())
		row := make([]int, W+1)
		for j := 1; j <= W; j++ {
			value, err := strconv.Atoi(line[j-1])
			if err != nil {
				fmt.Println("error:", err)
				return
			}
			row[j] = value
		}
		X[i] = row
	}

	// 累積和
	Z := make([][]int, H+1)
	Z[0] = make([]int, W+1)
	for i := 1; i <= H; i++ {
		Z[i] = make([]int, W+1)
		for j := 1; j <= W; j++ {
			Z[i][j] = X[i][j] + Z[i-1][j] + Z[i][j-1] - Z[i-1][j-1]
		}
	}

	sc.Scan()
	line := strings.Fields(sc.Text())
	Q, _ := strconv.Atoi(line[0])

	results := make([]int, Q)
	for i := 0; i < Q; i++ {
		sc.Scan()
		line := strings.Fields(sc.Text())
		A, _ := strconv.Atoi(line[0])
		B, _ := strconv.Atoi(line[1])
		C, _ := strconv.Atoi(line[2])
		D, _ := strconv.Atoi(line[3])
		results[i] = Z[C][D] + Z[A-1][B-1] - Z[A-1][D] - Z[C][B-1]
	}

	writer := bufio.NewWriter(os.Stdout)
	for _, result := range results {
		fmt.Fprintln(writer, result)
	}
	writer.Flush()
}
