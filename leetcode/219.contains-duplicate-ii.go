/*
 * @lc app=leetcode id=219 lang=golang
 *
 * [219] Contains Duplicate II
 */
package leetcode

// @lc code=start
func containsNearbyDuplicate(nums []int, k int) bool {
	if k <= 0 {
		return false
	}

	store := map[int]int{}
	for i, n := range nums {
		if _, ok := store[n]; !ok {
			store[n] = i
			continue
		}

		curr := store[n]
		if abs(i-curr) <= k {
			return true
		} else {
			store[n] = i
		}
	}

	return false
}

func abs(i int) int {
	if i < 0 {
		return -i
	}

	return i
}

// @lc code=end
