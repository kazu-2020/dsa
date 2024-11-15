/*
 * @lc app=leetcode id=643 lang=golang
 *
 * [643] Maximum Average Subarray I
 */

// @lc code=start
package leetcode

func findMaxAverage(nums []int, k int) float64 {
	max_sum := 0
	for i := 0; i < k; i++ {
		max_sum += nums[i]
	}
	curr_sum := max_sum

	for i := k; i < len(nums); i++ {
		curr_sum += nums[i] - nums[i-k]
		if curr_sum > max_sum {
			max_sum = curr_sum
		}
	}

	return float64(max_sum) / float64(k)
}
// @lc code=end
