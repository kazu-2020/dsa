package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Buffer(make([]byte, 1000000), 1000000)
	var N, K int
	sc.Scan()
	fmt.Sscanf(sc.Text(), "%d %d", &N, &K)

	sc.Scan()
	S := sc.Text()

	numON := 0
	for _, v := range S {
		if v == '1' {
			numON++
		}
	}

	if numON%2 == K%2 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
