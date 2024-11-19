package main

import "fmt"

func main() {
	var H, W int
	fmt.Scan(&H, &W)

	if H == 1 || W == 1 {
		fmt.Println(1)
		return
	}

	ans := H * W / 2
	if H*W%2 == 1 {
		ans++
	}

	fmt.Println(ans)
}
