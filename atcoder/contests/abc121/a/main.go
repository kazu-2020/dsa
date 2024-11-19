package main

import "fmt"

func main() {
	var H, W, h, w int
	fmt.Scan(&H, &W)
	fmt.Scan(&h, &w)

	ans := (H - h) * (W - w)
	fmt.Println(ans)
}
