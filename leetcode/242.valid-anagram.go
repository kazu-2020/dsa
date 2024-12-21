/*
 * @lc app=leetcode id=242 lang=golang
 *
 * [242] Valid Anagram
 */
package leetcode

// @lc code=start
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	store := make([]int, 26)
	for i := range len(s) {
		store[s[i]-'a']++
		store[t[i]-'a']--
	}

	for _, v := range store {
		if v != 0 {
			return false
		}
	}

	return true
}

// @lc code=end
