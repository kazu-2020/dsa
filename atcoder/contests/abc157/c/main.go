package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)

	var N, M int
	sc.Scan()
	fmt.Sscanf(sc.Text(), "%d %d", &N, &M)

	S, C := make([]int, M), make([]int, M)
	for i := 0; i < M; i++ {
		sc.Scan()
		fmt.Sscanf(sc.Text(), "%d %d", &S[i], &C[i])
	}

	for i := 0; i < 1000; i++ {
		if len(fmt.Sprintf("%d", i)) != N {
			continue
		}

		strI := fmt.Sprintf("%0*d", N, i)
		valid := true
		for j := 0; j < M; j++ {
			if int(strI[S[j]-1]-'0') != C[j] {
				valid = false
				break
			}
		}

		if valid {
			fmt.Println(i)
			return
		}
	}

	fmt.Println(-1)
}
