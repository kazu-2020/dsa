package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	sc.Scan()

	line := strings.Fields(sc.Text())
	prices := make([]int, len(line))
	for i, input := range line {
		prices[i], _ = strconv.Atoi(input)
	}

	for i := 0; i < len(prices); i++ {
		seenPrices := make(map[int]bool)
		for j := i + 1; j < len(prices); j++ {
			target := 1000 - prices[i] - prices[j]
			if seenPrices[target] {
				fmt.Println("Yes")
				return
			}

			seenPrices[prices[j]] = true
		}
	}

	fmt.Println("No")
}
