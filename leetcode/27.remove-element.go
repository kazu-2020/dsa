package leetcode

/*
 * @lc app=leetcode id=27 lang=golang
 *
 * [27] Remove Element
 */

// nums = [3,2,2,3], val = 3
// expected: 2, nums = [2,2,_,_]

// @lc code=start
func removeElement(nums []int, val int) int {
	k := 0

	for _, n := range nums {
		if n != val {
			nums[k] = n
			k++
		}
	}

	return k
}
// @lc code=end
