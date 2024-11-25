package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)

	ans := 1
	most_count := 0
	for i := 1; i <= N; i++ {
		curr := i
		count := 0
		for curr%2 == 0 {
			curr /= 2
			count++
		}

		if most_count < count {
			ans = i
			most_count = count
		}
	}

	fmt.Println(ans)
}
