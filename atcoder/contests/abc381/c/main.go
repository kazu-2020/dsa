package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var N int
	fmt.Scan(&N)
	sc := bufio.NewScanner(os.Stdin)
	sc.Buffer(make([]byte, 100000000), 100000000)
	sc.Scan()
	S := sc.Text()

	ans := 1
	for i := 0; i < N; i++ {
		curr := S[i]
		if curr != '/' {
			continue
		}

		d := 0
		for {
			l := i - (d + 1)
			r := i + (d + 1)
			if l < 0 || N <= l {
				break
			}
			if r < 0 || N <= r {
				break
			}
			if S[l] != '1' {
				break
			}
			if S[r] != '2' {
				break
			}
			d++
		}

		ans = max(ans, 1+d*2)
	}

	fmt.Println(ans)
}

func max(a, b int) int {
	if a < b {
		return b
	}

	return a
}
