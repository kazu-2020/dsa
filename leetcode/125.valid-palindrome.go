/*
 * @lc app=leetcode id=125 lang=golang
 *
 * [125] Valid Palindrome
 */

package leetcode

import "strings"

// @lc code=start

func isPalindrome(s string) bool {
	left, right := 0, len(s)-1

	for left < right {
		for left < right && !isAlphaNumeric(s[left]) {
			left++
		}
		for left < right && !isAlphaNumeric(s[right]) {
			right--
		}

		if !strings.EqualFold(string(s[left]), string(s[right])) {
			return false
		}

		left++
		right--
	}

	return true
}

func isAlphaNumeric(r byte) bool {
	return ('a' <= r && r <= 'z') || ('A' <= r && r <= 'Z') || ('0' <= r && r <= '9')
}

// @lc code=end
