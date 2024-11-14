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
	// あたり、はずれの数列
	sc.Scan()
	line := strings.Fields(sc.Text())
	A := make([]int, N+1)
	for i := 1; i <= N; i++ {
		A[i], _ = strconv.Atoi(line[i-1])
	}

	// あたり、はずれの数列の累積和
	hit_count := make([]int, N+1)
	los_count := make([]int, N+1)
	hit_count[0] = 0
	los_count[0] = 0
	for i := 1; i <= N; i++ {
		current := A[i]
		if current == 1 {
			hit_count[i] = hit_count[i-1] + 1
			los_count[i] = los_count[i-1]
		} else {
			hit_count[i] = hit_count[i-1]
			los_count[i] = los_count[i-1] + 1
		}
	}

	sc.Scan()
	Q, _ := strconv.Atoi(sc.Text())
	for i := 1; i <= Q; i++ {
		sc.Scan()
		line := strings.Fields(sc.Text())
		L, _ := strconv.Atoi(line[0])
		R, _ := strconv.Atoi(line[1])
		hit := hit_count[R] - hit_count[L-1]
		los := los_count[R] - los_count[L-1]

		switch {
		case hit > los:
			fmt.Println("win")
		case hit == los:
			fmt.Println("draw")
		default:
			fmt.Println("lose")
		}
	}
}
