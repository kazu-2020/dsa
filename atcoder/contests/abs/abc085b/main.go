package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)

	d := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&d[i])
	}

	store := make(map[int]bool)
	uniq := []int{}
	// 重複を取り除く
	for _, v := range d {
		if _, ok := store[v]; !ok {
			store[v] = true
			uniq = append(uniq, v)
		}
	}

	fmt.Println(len(uniq))
}
