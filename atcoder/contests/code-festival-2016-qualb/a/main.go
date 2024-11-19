package main

import "fmt"

const COLLECT = "CODEFESTIVAL2016"

func main() {
	var S string
	fmt.Scan(&S)

	ans := 0

	for i := range S {
		if S[i] != COLLECT[i] {
			ans++
		}
	}

	fmt.Println(ans)
}
