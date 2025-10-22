package code

/*
 * @lc app=leetcode id=704 lang=golang
 *
 * [704] Binary Search
 */

// @lc code=start
func search(nums []int, target int) int {
	// nums = [-1,0,3,5,9,12], target = 9
	// n = 6
	n := len(nums)

	// left = 0, right = 5
	left, right := 0, n-1

	// Time complexity: O(log n)
	// Space complexity: O(1)
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

// @lc code=end
