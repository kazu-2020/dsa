/*
 * @lc app=leetcode id=953 lang=golang
 *
 * [953] Verifying an Alien Dictionary
 */
package leetcode

// @lc code=start
func isAlienSorted(words []string, order string) bool {
	if len(words) == 1 {
		return true
	}

	// 辞書(order)の各文字が何番目かを管理
	store := map[rune]int{}
	for i, v := range order {
		store[v] = i
	}

	for i := 0; i < len(words)-1; i++ {
		curr := words[i]
		next := words[i+1]

		minLength := min(len(curr), len(next))
		if curr[:minLength] == next[:minLength] && len(curr) > len(next) {
			return false
		}

		pos := 0
		for pos < minLength {
			if store[rune(curr[pos])] < store[rune(next[pos])] {
				break
			}
			if store[rune(curr[pos])] > store[rune(next[pos])] {
				return false
			}

			pos++
		}
	}

	return true
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

// @lc code=end
