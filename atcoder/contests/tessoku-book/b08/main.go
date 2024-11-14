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

	// 座標 X, Y は 1 以上 1500 以下
	grid_size := 1501
	grid := make([][]int, grid_size)
	for i := 0; i < grid_size; i++ {
		grid[i] = make([]int, grid_size)
	}
	// 各座標を grid にマッピング
	for i := 1; i <= N; i++ {
		sc.Scan()
		line := sc.Text()
		X, _ := strconv.Atoi(strings.Fields(line)[0])
		Y, _ := strconv.Atoi(strings.Fields(line)[1])

		grid[X][Y] += 1
	}
	// 累積和
	sum := make([][]int, grid_size)
	for i := 0; i < grid_size; i++ {
		sum[i] = make([]int, grid_size)
	}
	for i := 1; i < grid_size; i++ {
		for j := 1; j < grid_size; j++ {
			sum[i][j] = grid[i][j] + sum[i-1][j] + sum[i][j-1] - sum[i-1][j-1]
		}
	}

	sc.Scan()
	line := sc.Text()
	Q, _ := strconv.Atoi(strings.Fields(line)[0])

	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	for i := 0; i < Q; i++ {
		sc.Scan()
		line := sc.Text()
		a, _ := strconv.Atoi(strings.Fields(line)[0])
		b, _ := strconv.Atoi(strings.Fields(line)[1])
		c, _ := strconv.Atoi(strings.Fields(line)[2])
		d, _ := strconv.Atoi(strings.Fields(line)[3])

		result := sum[c][d] + sum[a-1][b-1] - sum[a-1][d] - sum[c][b-1]
		fmt.Fprintln(writer, result)
	}

}
