package code

/*
 * @lc app=leetcode id=153 lang=golang
 *
 * [153] Find Minimum in Rotated Sorted Array
 */

// @lc code=start
func findMin(nums []int) int {
	n := len(nums)

	left, right := 0, n-1
	min := nums[0]

	for left <= right {
		mid := left + (right-left)/2

		if nums[mid] >= nums[left] {
			if nums[left] < min {
				min = nums[left]
			}
			left = mid + 1
		} else {
			if nums[mid] < min {
				min = nums[mid]
			}
			right = mid - 1
		}
	}
	return min
}

// @lc code=end
