/*
 * @lc app=leetcode id=1002 lang=golang
 *
 * [1002] Find Common Characters
 */
package leetcode

import (
	"math"
)

// @lc code=start

func commonChars(words []string) []string {
	store := [26]int{}
	for i := 0; i < 26; i++ {
		store[i] = math.MaxInt64
	}

	for _, word := range words {
		curr := [26]int{}
		// 単語内にアルファベットがそれぞれいくつあるかを計算
		for _, chr := range word {
			curr[chr-'a']++
		}

		for i := 0; i < 26; i++ {
			if curr[i] < store[i] {
				store[i] = curr[i]
			}
		}
	}

	ans := []string{}
	for i, _ := range store {
		for 0 < store[i] {
			ans = append(ans, string('a'+i))
			store[i]--
		}
	}

	return ans
}

// @lc code=end

// func commonChars(words []string) []string {
// 	// 最小頻度を記録する(a-z)
// 	store := make([]int, 26)
// 	for i := 0; i < 26; i++ {
// 		store[i] = math.MaxInt64
// 	}

// 	for _, word := range words {
// 		curr := make([]int, 26)
// 		for _, char := range word {
// 			curr[int(char-'a')]++
// 		}

// 		for i := 0; i < 26; i++ {
// 			if curr[i] < store[i] {
// 				store[i] = curr[i]
// 			}
// 		}
// 	}

// 	ans := []string{}
// 	for i := 0; i < 26; i++ {
// 		for store[i] > 0 {
// 			ans = append(ans, string('a'+i))
// 			store[i]--
// 		}
// 	}

// 	return ans
// }
