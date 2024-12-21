/*
 * @lc app=leetcode id=128 lang=golang
 *
 * [128] Longest Consecutive Sequence
 */
package leetcode

// @lc code=start
func longestConsecutive(nums []int) int {
	store := map[int]bool{}
	for _, v := range nums {
		store[v] = true
	}

	best := 0
	for num := range store {
		// 1 つ前の値があれば、開始地点ではない
		if store[num-1] {
			continue
		}

		curr := num
		streak := 1
		// 連続する値が続くまでループする
		for store[curr+1] {
			curr++
			streak++
		}

		if streak > best {
			best = streak
		}
	}

	return best
}

// @lc code=end
