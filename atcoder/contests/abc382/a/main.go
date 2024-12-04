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
	s := 0
	for _, v := range sc.Text() {
		if v == '.' {
			s++
		}
	}

	fmt.Println(s + D)
}
