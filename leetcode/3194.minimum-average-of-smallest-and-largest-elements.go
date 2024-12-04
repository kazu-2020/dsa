/*
 * @lc app=leetcode id=3194 lang=golang
 *
 * [3194] Minimum Average of Smallest and Largest Elements
 */

package leetcode

import (
	"math"
	"slices"
)

// @lc code=start
func minimumAverage(nums []int) float64 {
	slices.Sort(nums)
	curr := math.MaxFloat64

	for i := 0; i < len(nums)/2; i++ {
		l := float64(nums[i])
		r := float64(nums[len(nums)-i-1])
		avg := (l + r) / 2.0

		if avg < curr {
			curr = avg
		}
	}

	return curr
}

// @lc code=end
