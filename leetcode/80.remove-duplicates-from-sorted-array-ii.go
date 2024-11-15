/*
 * @lc app=leetcode id=80 lang=golang
 *
 * [80] Remove Duplicates from Sorted Array II
 */

// @lc code=start
package leetcode

func removeDuplicates(nums []int) int {
	n := len(nums)
	if n <= 2 {
		return n
	}

	i := 2
	for j := 2; j < n; j++ {
		if nums[j] != nums[i-2] {
			nums[i] = nums[j]
			i++
		}
	}

	return i
}
// @lc code=end

// func removeDuplicates(nums []int) int {
// 	store := make(map[int]int)
// 	curr := 0

// 	for i := 0; i < len(nums); i++ {
// 		if store[nums[i]] < 2 {
// 			store[nums[i]]++
// 			nums[curr] = nums[i]
// 			curr++
// 		}
// 	}

// 	return curr
// }
