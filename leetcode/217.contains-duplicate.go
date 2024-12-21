/*
 * @lc app=leetcode id=217 lang=golang
 *
 * [217] Contains Duplicate
 */
package leetcode

// @lc code=start
func containsDuplicate(nums []int) bool {
	store := map[int]int{}
	for _, v := range nums {
		store[v]++

		if store[v] >= 2 {
			return true
		}
	}

	return false
}

// @lc code=end
