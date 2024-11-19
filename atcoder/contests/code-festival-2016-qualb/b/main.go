package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Buffer(make([]byte, 10000), 1000000)

	var N, A, B int
	sc.Scan()
	fmt.Sscanf(sc.Text(), "%d %d %d", &N, &A, &B)

	sc.Scan()
	S := sc.Text()

	var curr, b_curr int
	for i := range S {
		if S[i] == 'a' && A+B > curr {
			curr++
			fmt.Println("Yes")
			continue
		}

		if S[i] == 'b' && A+B > curr && B > b_curr {
			curr++
			b_curr++
			fmt.Println("Yes")
			continue
		}

		fmt.Println("No")
	}
}
