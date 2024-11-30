package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	var N, K int
	sc.Scan()
	fmt.Sscanf(sc.Text(), "%d %d", &N, &K)

	if K >= 2*N-2 && K%2 == 0 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
