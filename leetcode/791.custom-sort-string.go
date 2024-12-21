/*
 * @lc app=leetcode id=791 lang=golang
 *
 * [791] Custom Sort String
 */
package leetcode

import (
	"cmp"
	"slices"
)

// @lc code=start
func customSortString(order string, s string) string {
	store := map[rune]int{}
	for i, v := range order {
		store[v] = i
	}

	arr := []rune(s)
	slices.SortFunc(arr, func(a, b rune) int {
		n := cmp.Compare(store[a], store[b])

		return n
	})

	return string(arr)
}

// @lc code=end
