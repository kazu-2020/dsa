package main

import "fmt"

func main() {
	var X string
	fmt.Scan(&X)

	same, step := true, true
	for i := 0; i < 3; i++ {
		if X[i] != X[i+1] {
			same = false
		}

		if (X[i]-'0'+1)%10 != (X[i+1]-'0')%10 {
			step = false
		}
	}

	if same || step {
		fmt.Println("Weak")
	} else {
		fmt.Println("Strong")
	}
}
