/*
 * @lc app=leetcode id=1886 lang=golang
 *
 * [1886] Determine Whether Matrix Can Be Obtained By Rotation
 */
package leetcode

/*
	90度回転 (i,j) => (j, n-i-1)
	1 2 3 .  7 4 1
	4 5 6 => 8 5 2
	7 8 9    9 6 3

	180度回転 (i,j) => (n-1-i, n-1-j)
	1 2 3 .  9 8 7
	4 5 6 => 6 5 4
	7 8 9    3 2 1

	270度回転 (i,j) => (n-1-j, i)
	1 2 3 .  3 6 9
	4 5 6 => 2 5 8
	7 8 9    1 4 7
*/

// @lc code=start
func findRotation(mat [][]int, target [][]int) bool {
	n := len(mat)
	results := make([]bool, 4)
	for i := range results {
		results[i] = true
	}

	for i, line := range mat {
		for j, val := range line {
			// 90 度
			if target[j][n-1-i] != val {
				results[0] = false
			}
			// 180 度
			if target[n-1-i][n-1-j] != val {
				results[1] = false
			}
			// 270 度
			if target[n-1-j][i] != val {
				results[2] = false
			}
			// 0 度
			if target[i][j] != val {
				results[3] = false
			}
		}
	}

	for _, val := range results {
		if val {
			return true
		}
	}

	return false
}

// @lc code=end
