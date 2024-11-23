package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)

	curr_t, curr_x, curr_y := 0, 0, 0
	for i := 1; i <= N; i++ {
		var t, x, y int
		fmt.Scan(&t, &x, &y)

		dt := abs(t - curr_t)
		distance := abs(curr_x-x) + abs(curr_y-y)
		fmt.Println(distance, dt)

		if distance > dt || (dt-distance)%2 != 0 {
			fmt.Println("No")
			return
		}

		curr_t, curr_x, curr_y = t, x, y
	}

	fmt.Println("Yes")
}

func abs(n int) int {
	if n < 0 {
		return -n
	}

	return n
}
