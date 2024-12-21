/*
 * @lc app=leetcode id=1572 lang=golang
 *
 * [1572] Matrix Diagonal Sum
 */

/*
	0, 1, 1, 1, 1, 0
	1, 0, 1, 1, 0, 1
	1, 1, 0, 0, 1, 1
	1, 0, 1, 1, 0, 1
	0, 1, 1, 1, 1, 0
*/

package leetcode

// @lc code=start
func diagonalSum(mat [][]int) int {
	n := len(mat)

	sum := 0
	left, right := 0, n-1
	for _, line := range mat {
		sum += line[left]
		if left != right {
			sum += line[right]
		}

		left++
		right--
	}

	return sum
}

// @lc code=end
