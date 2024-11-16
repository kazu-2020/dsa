package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	var N int
	fmt.Sscanf(sc.Text(), "%d", &N)

	prevT, prevX, prevY := 0, 0, 0
	for i := 0; i < N; i++ {
		sc.Scan()
		var T, X, Y int
		fmt.Sscanf(sc.Text(), "%d %d %d", &T, &X, &Y)

		dt := T - prevT
		distance := abs(X-prevX) + abs(Y-prevY)

		if distance > dt || (dt-distance)%2 != 0 {
			fmt.Println("No")
			return
		}

		prevT, prevX, prevY = T, X, Y
	}

	fmt.Println("Yes")
}

func abs(a int) int {
	if a < 0 {
		return -a
	}

	return a
}
