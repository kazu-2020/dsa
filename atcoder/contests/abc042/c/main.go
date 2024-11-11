package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
    scanner := bufio.NewScanner(os.Stdin)

    // NとKの読み込み
    scanner.Scan()
    firstLine := strings.Fields(scanner.Text())
    N, _ := strconv.Atoi(firstLine[0])
    K, _ := strconv.Atoi(firstLine[1])

    // 嫌いな数字をセットに格納
		// 嫌いな数字が 1,3,4 の場合、
		// { '1': true, '3': true, '4': true } というセットになる
    dislikedSet := make(map[rune]bool)
    if K > 0 {
        scanner.Scan()
        dislikedNumbers := strings.Fields(scanner.Text())
        for _, d := range dislikedNumbers {
            dislikedSet[rune(d[0])] = true
        }
    }

    // 嫌いな数字がない場合、そのまま N を出力
    if K == 0 {
        fmt.Println(N)
        return
    }

    // N以上の数で嫌いな数字を含まない最小の金額を探す
    for i := N; ; i++ {
        price := strconv.Itoa(i)
        if isAcceptable(price, dislikedSet) {
            fmt.Println(price)
            break
        }
    }
}

// 金額が嫌いな数字を含むかどうかをチェックする関数
func isAcceptable(price string, dislikedSet map[rune]bool) bool {
    for _, digit := range price {
        if dislikedSet[digit] {
            return false
        }
    }
    return true
}
