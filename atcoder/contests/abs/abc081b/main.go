package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var N int
	fmt.Scan(&N)

	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	line := strings.Fields(sc.Text())

	nums := make([]int, N)
	for i := range line {
		nums[i], _ = strconv.Atoi(line[i])
	}

	ans := 0
	flag := true

	for flag {
		for i, num := range nums {
			if num % 2 == 0 {
				nums[i] = num / 2
			} else {
				flag = false
				break
			}
		}

		if flag {
			ans++
		}
	}

	fmt.Println(ans)
}
