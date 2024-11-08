package leetcode

/*
 * @lc app=leetcode id=1480 lang=golang
 *
 * [1480] Running Sum of 1d Array
 */

// @lc code=start
func runningSum(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		if i > 0 {
			nums[i] += nums[i-1]
		}
	}

	return nums
}
// @lc code=end
