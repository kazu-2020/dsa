package leetcode

/*
 * @lc app=leetcode id=344 lang=golang
 *
 * [344] Reverse String
 */

// @lc code=start
func reverseString(s []byte)  {
    start  := 0
		s_nums := len(s) - 1

		for start < s_nums {
			s[start], s[s_nums] = s[s_nums], s[start]

			start++
			s_nums--
		}
}
// @lc code=end
