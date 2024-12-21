/*
 * @lc app=leetcode id=3005 lang=golang
 *
 * [3005] Count Elements With Maximum Frequency
 */
package leetcode

// @lc code=start
func maxFrequencyElements(nums []int) int {
	store := map[int]int{}
	max := 0 // 最も多い登場回数
	for _, v := range nums {
		store[v]++
		if store[v] > max {
			max = store[v]
		}
	}

	sum := 0
	for _, v := range store {
		if v == max {
			sum += v
		}
	}

	return sum
}

// @lc code=end
