/*
 * @lc app=leetcode id=1512 lang=golang
 *
 * [1512] Number of Good Pairs
 */
package leetcode

// @lc code=start
func numIdenticalPairs(nums []int) int {
	store := map[int]int{}

	for _, num := range nums {
		store[num]++
	}

	ans := 0
	for _, val := range store {
		ans += val * (val - 1) / 2
	}

	return ans
}

// @lc code=end

// func numIdenticalPairs(nums []int) int {
// 	ans := 0
// 	for i := 0; i < len(nums); i++ {
// 		for j := i + 1; j < len(nums); j++ {
// 			if nums[i] == nums[j] {
// 				ans++
// 			}
// 		}
// 	}

// 	return ans
// }
