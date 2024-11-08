package leetcode

/*
 * @lc app=leetcode id=88 lang=golang
 *
 * [88] Merge Sorted Array
 */

// nums1 = [1,2,3,0,0,0], m = 3
// nums2 = [2,5,6],       n = 3

// @lc code=start
func merge(nums1 []int, m int, nums2 []int, n int)  {
	i, j, k := m-1, n-1, m+n-1

	for j >= 0 {
		if i >= 0 && nums1[i] > nums2[j] {
			nums1[k] = nums1[i]
			i--
		} else {
			nums1[k] = nums2[j]
			j--
		}

		k--
	}
}
// @lc code=end

// func merge(nums1 []int, m int, nums2 []int, n int)  {
// 	i := m - 1
// 	j := n - 1
// 	k := m + n - 1

// 	for i >= 0 && j >= 0 {
// 		if nums1[i] > nums2[j] {
// 			nums1[k] = nums1[i]
// 			i--
// 		} else {
// 			nums1[k] = nums2[j]
// 			j--
// 		}

// 		k--
// 	}

// 	for j >= 0 {
// 		nums1[k] = nums2[j]
// 		j--
// 		k--
// 	}

// }