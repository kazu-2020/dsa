package main

import "fmt"

func main() {
	var S string
	fmt.Scan(&S)

	flag := true
	size := len(S)

	if size%2 != 0 {
		flag = false
	}

	if flag {
		for i := 0; i < size/2; i++ {
			if S[2*i] != S[2*i+1] {
				flag = false
				break
			}
		}
	}

	if flag {
		store := map[byte]int{}
		for i := 0; i < size; i++ {
			store[S[i]]++
		}
		for _, v := range store {
			if v != 2 {
				flag = false
			}
		}
	}

	if flag {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
