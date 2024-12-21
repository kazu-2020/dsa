/*
 * @lc app=leetcode id=49 lang=golang
 *
 * [49] Group Anagrams
 */
package leetcode

// @lc code=start
func groupAnagrams(strs []string) [][]string {
	store := make(map[[26]int][]string)
	for _, str := range strs {
		curr := [26]int{}
		for _, chr := range str {
			curr[chr-'a']++
		}

		val, ok := store[curr]
		if ok {
			store[curr] = append(val, str)
		} else {
			store[curr] = []string{str}
		}
	}

	ans := [][]string{}
	for _, val := range store {
		ans = append(ans, val)
	}

	return ans
}

// @lc code=end
