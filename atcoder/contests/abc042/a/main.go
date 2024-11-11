package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()

	inputs := strings.Split(sc.Text(), " ")
	sort.Strings(inputs)

	if inputs[0] == "5" && inputs[1] == "5" && inputs[2] == "7" {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
