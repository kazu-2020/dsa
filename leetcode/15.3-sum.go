/*
 * @lc app=leetcode id=15 lang=golang
 *
 * [15] 3Sum
 */
package leetcode

import (
	"slices"
)

// @lc code=start
func threeSum(nums []int) [][]int {
	ans := [][]int{}

	slices.Sort(nums)
	if len(nums) < 3 {
		return ans
	}
	if nums[0] > 0 {
		return ans
	}

	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		l := i + 1
		r := len(nums) - 1

		for l < r {
			total := nums[i] + nums[l] + nums[r]

			if total > 0 {
				r--
			} else if total < 0 {
				l++
			} else {
				valid := []int{nums[i], nums[l], nums[r]}
				ans = append(ans, valid)
				l++

				for nums[l] == nums[l-1] && l < r {
					l++
				}
			}
		}
	}

	return ans
}

// @lc code=end
