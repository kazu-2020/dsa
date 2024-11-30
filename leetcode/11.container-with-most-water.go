/*
 * @lc app=leetcode id=11 lang=golang
 *
 * [11] Container With Most Water
 */
package leetcode

// @lc code=start
func maxArea(height []int) int {
	l, r := 0, len(height)-1

	ans := 0
	for l < r {
		s := min(height[l], height[r]) * (r - l)
		if s > ans {
			ans = s
		}

		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}

	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

// @lc code=end
