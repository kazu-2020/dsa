package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)

	var N, D int
	sc.Scan()
	fmt.Sscanf(sc.Text(), "%d %d", &N, &D)

	sc.Scan()
	S := []rune(sc.Text())
	for i := len(S) - 1; i >= 0; i-- {
		if S[i] == '@' {
			S[i] = '.'
			D--
		}

		if D == 0 {
			break
		}
	}

	fmt.Println(string(S))
}
