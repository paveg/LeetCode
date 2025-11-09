package code

/*
 * @lc app=leetcode id=53 lang=golang
 *
 * [53] Maximum Subarray
 */

// @lc code=start
func maxSubArray(nums []int) int {
	curSum, maxSum := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		curSum = max(nums[i], curSum+nums[i])
		maxSum = max(maxSum, curSum)
	}
	return maxSum
}

// @lc code=end
