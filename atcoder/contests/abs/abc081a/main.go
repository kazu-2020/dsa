package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)

	ans := 0
	for i := range s {
		if s[i] == '1' {
			ans++
		}
	}

	fmt.Println(ans)
}
