package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)

	var N, M int // N: 店の件数, M: 買いたい本数
	sc.Scan()
	fmt.Sscanf(sc.Text(), "%d %d", &N, &M)

	shops := make([]Shop, N)
	for i := 0; i < N; i++ {
		sc.Scan()
		line := strings.Fields(sc.Text())

		price, _ := strconv.Atoi(line[0])
		stock, _ := strconv.Atoi(line[1])
		shops[i] = Shop{price, stock}
	}

	// 値段が安い順に店を並び替え
	sort.Slice(shops, func(i, j int) bool {
		return shops[i].price < shops[j].price
	})

	cost := 0
	for _, shop := range shops {
		if M <= 0 {
			break
		}

		buy := min(shop.stock, M)
		cost += shop.price * buy
		M -= buy
	}

	fmt.Println(cost)
}

type Shop struct {
	price int
	stock int
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
