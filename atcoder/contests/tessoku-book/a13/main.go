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
	sc.Buffer(make([]byte, 10000000), 10000000)

	var N, K int
	sc.Scan()
	fmt.Sscanf(sc.Text(), "%d %d", &N, &K)

	A := make([]int, N+1)
	sc.Scan()
	line := strings.Fields(sc.Text())
	for i := 1; i <= N; i++ {
		A[i], _ = strconv.Atoi(line[i-1])
	}

	R := make([]int, N+1)
	for i := 1; i <= N; i++ {
		if i == 1 {
			R[i] = 1
		} else {
			R[i] = R[i-1]
		}

		for {
			if R[i] < N && A[R[i]+1]-A[i] <= K {
				R[i] += 1
				continue
			}

			break
		}
	}

	ans := 0
	for i := 1; i <= N; i++ {
		ans += R[i] - i
	}

	fmt.Println(ans)
}
