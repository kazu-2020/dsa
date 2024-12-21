/*
 * @lc app=leetcode id=2191 lang=golang
 *
 * [2191] Sort the Jumbled Numbers
 */
package leetcode

import (
	"cmp"
	"slices"
	"strconv"
)

// @lc code=start
func sortJumbled(mapping []int, nums []int) []int {
	slices.SortStableFunc(nums, func(a, b int) int {
		i := mappedNum(a, mapping)
		j := mappedNum(b, mapping)

		return cmp.Compare(i, j)
	})

	return nums
}

func mappedNum(num int, mapping []int) int {
	res := 0
	str := strconv.Itoa(num)
	for _, chr := range str {
		res = res*10 + mapping[chr-'0']
	}

	return res
}

// @lc code=end
