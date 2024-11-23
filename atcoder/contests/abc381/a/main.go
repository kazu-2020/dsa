package main

import "fmt"

func main() {
	var N int
	var S string
	fmt.Scan(&N)
	fmt.Scan(&S)

	if N%2 == 0 {
		no()
		return
	}

	for i := 0; i <= (N-1)/2-1; i++ {
		if S[i] != '1' {
			no()
			return
		}
	}

	if S[(N-1)/2] != '/' {
		no()
		return
	}

	for i := (N-1)/2 + 1; i < N; i++ {
		if S[i] != '2' {
			no()
			return
		}
	}

	fmt.Println("Yes")
}

func no() {
	fmt.Println("No")
}
