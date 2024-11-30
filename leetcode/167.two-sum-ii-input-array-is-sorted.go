/*
 * @lc app=leetcode id=167 lang=golang
 *
 * [167] Two Sum II - Input Array Is Sorted
 */
package leetcode

// @lc code=start
func twoSum(numbers []int, target int) []int {
	l, r := 0, len(numbers)-1
	ans := make([]int, 2)
	for l < r {
		s := numbers[l] + numbers[r]
		if s == target {
			ans[0] = l + 1
			ans[1] = r + 1
			break
		} else if s < target {
			l++
		} else {
			r--
		}
	}

	return ans
}

// @lc code=end
