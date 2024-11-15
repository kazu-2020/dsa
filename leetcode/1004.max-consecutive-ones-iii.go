/*
 * @lc app=leetcode id=1004 lang=golang
 *
 * [1004] Max Consecutive Ones III
 */

// @lc code=start
package leetcode

func longestOnes(nums []int, k int) int {
	var left, curr, ans int

	for right := 0; right < len(nums); right++ {
		if nums[right] == 0 {
			curr++
		}

		for curr > k {
			if nums[left] == 0 {
				curr--
			}
			left++
		}

		t := right - left + 1
		if t > ans {
			ans = t
		}
	}

	return ans
}
// @lc code=end
