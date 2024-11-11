package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	first_input := strings.Fields(sc.Text())
	x := first_input[1]

	sc.Scan()
	string_list := strings.Fields(sc.Text())

	// slices は 1.20.6 ではサポートされていない。。
	// result := slices.Contains(string_list, x)
	// if result {
	// 	fmt.Println("Yes")
	// } else {
	// 	fmt.Println("No")
	// }

	for _, s := range string_list {
		if s == x {
			fmt.Println("Yes")
			return
		}
	}

	fmt.Println("No")
}
