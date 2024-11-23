package main

import (
	"fmt"
)

func main() {
	grid := make([][]bool, 4)
	for i := 1; i <= 3; i++ {
		grid[i] = make([]bool, 4)
	}

	bingo := map[int]Position{}
	for i := 1; i <= 3; i++ {
		var A1, A2, A3 int
		fmt.Scan(&A1, &A2, &A3)

		bingo[A1] = Position{i: i, j: 1}
		bingo[A2] = Position{i: i, j: 2}
		bingo[A3] = Position{i: i, j: 3}
	}

	var N int
	fmt.Scan(&N)
	for i := 1; i <= N; i++ {
		var b int
		fmt.Scan(&b)

		if val, ok := bingo[b]; ok {
			grid[val.i][val.j] = true
		}
	}

	// 横に揃っているか?
	for i := 1; i <= 3; i++ {
		a := grid[i][1]
		b := grid[i][2]
		c := grid[i][3]

		if a && b && c {
			fmt.Println("Yes")
			return
		}
	}

	// 縦に揃っているか?
	for i := 1; i <= 3; i++ {
		a := grid[1][i]
		b := grid[2][i]
		c := grid[3][i]

		if a && b && c {
			fmt.Println("Yes")
			return
		}
	}

	// 斜めに揃っているか?
	if grid[1][1] && grid[2][2] && grid[3][3] {
		fmt.Println("Yes")
		return
	}
	if grid[1][3] && grid[2][2] && grid[3][1] {
		fmt.Println("Yes")
		return
	}

	fmt.Println("No")
}

type Position struct {
	i int
	j int
}
