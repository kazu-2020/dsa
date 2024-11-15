/*
 * @lc app=leetcode id=121 lang=golang
 *
 * [121] Best Time to Buy and Sell Stock
 */

// @lc code=start
package leetcode

import "math"

func maxProfit(prices []int) int {
	minPrice := math.MaxInt32
	benefit := 0

	for i := 0; i < len(prices); i++ {
		if prices[i] < minPrice {
			minPrice = prices[i]
		} else if prices[i] - minPrice > benefit {
			benefit = prices[i] - minPrice
		}
	}

	return benefit
}
// @lc code=end
