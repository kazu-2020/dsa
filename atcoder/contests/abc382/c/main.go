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
	sc.Buffer(make([]byte, 100000000), 100000000)

	var N, M int
	sc.Scan()
	fmt.Sscanf(sc.Text(), "%d %d", &N, &M)

	sc.Scan()
	line := strings.Fields(sc.Text())

	ids := make([]int, 200010)
	for i := range ids {
		ids[i] = -1
	}
	r := 200010
	for i := 0; i < N; i++ {
		level, _ := strconv.Atoi(line[i])

		for r > level {
			r--
			ids[r] = i + 1
		}
	}

	sc.Scan()
	line = strings.Fields(sc.Text())
	for i := 0; i < M; i++ {
		b, _ := strconv.Atoi(line[i])

		fmt.Println(ids[b])
	}
}
