package main

import "fmt"

func main() {
	var S, T string
	fmt.Scan(&S, &T)

	ans := 0
	for i := 0; i < 3; i++ {
		if S[i] == T[i] {
			ans++
		}
	}

	fmt.Println(ans)
}
