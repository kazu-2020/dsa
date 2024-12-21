/*
 * @lc app=leetcode id=1 lang=golang
 *
 * [1] Two Sum
 */
package leetcode

// @lc code=start
func twoSum(nums []int, target int) []int {
	store := map[int]int{}

	for i, v := range nums {
		t := target - v

		if val, ok := store[t]; ok {
			return []int{val, i}
		}

		store[v] = i
	}

	return []int{}
}

// @lc code=end
