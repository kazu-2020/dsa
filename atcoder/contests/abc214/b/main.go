package main

import "fmt"

func main() {
	var S, T int
	fmt.Scan(&S, &T)

	ans := 0
	for a := 0; a <= 100; a++ {
		for b := 0; b <= 100; b++ {
			for c := 0; c <= 100; c++ {
				if a+b+c <= S && a*b*c <= T {
					ans++
				} else {
					break
				}
			}
		}
	}
	fmt.Println(ans)
}
