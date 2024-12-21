/*
 * @lc app=leetcode id=179 lang=golang
 *
 * [179] Largest Number
 */
package leetcode

import (
	"slices"
	"strconv"
	"strings"
)

// @lc code=start
func largestNumber(nums []int) string {
	strs := make([]string, len(nums))
	for i, v := range nums {
		strs[i] = strconv.Itoa(v)
	}

	slices.SortStableFunc(strs, func(a, b string) int {
		i := a + b
		j := b + a

		if i < j {
			return 1
		} else {
			return -1
		}
	})

	ans := strings.Join(strs, "")
	for _, v := range ans {
		if v != '0' {
			return ans
		}
	}

	return "0"
}

// @lc code=end
