package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Buffer(make([]byte, 1000000), 1000000)
	sc.Scan()
	S := sc.Text()

	for {
		switch {
			case strings.HasSuffix(S, "dream"):
				S = strings.TrimSuffix(S, "dream")
			case strings.HasSuffix(S, "dreamer"):
				S = strings.TrimSuffix(S, "dreamer")
			case strings.HasSuffix(S, "erase"):
				S = strings.TrimSuffix(S, "erase")
			case strings.HasSuffix(S, "eraser"):
				S = strings.TrimSuffix(S, "eraser")
			case len(S) == 0:
				fmt.Println("YES")
				return
			default:
				fmt.Println("NO")
				return
		}
	}
}
