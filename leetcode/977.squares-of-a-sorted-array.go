package leetcode

/*
 * @lc app=leetcode id=977 lang=golang
 *
 * [977] Squares of a Sorted Array
 */

// @lc code=start
func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func sortedSquares(nums []int) []int {
	result := make([]int, len(nums))
	left, right := 0, len(nums)-1

	for i := len(nums) - 1; i >= 0; i-- {
		square := 0

		if abs(nums[left]) > abs(nums[right]) {
			square = nums[left]
			left++
		} else {
			square = nums[right]
			right--
		}

		result[i] = square * square
	}

	return result
}

// @lc code=end

// func sortedSquares(nums []int) []int {
// 	for i := range nums {
// 		nums[i] *= nums[i]
// 	}
// 	sort.Ints(nums)

// 	return nums
// }
