package code

/*
 * @lc app=leetcode id=198 lang=golang
 *
 * [198] House Robber
 */

// @lc code=start
func rob(nums []int) int {
	n := len(nums)
	// 1 <= n <= 100, so if n == 1, we can directly return nums[0]
	if n == 1 {
		return nums[0]
	}

	dp := make([]int, n)
	// Example1: nums = [1,2,3,1]
	// dp[0] = 1, dp[1] = max(1,2) = 2, dp[2] = max(2,1+3) = 4, dp[3] = max(4,2+1) = 4
	// Example2: nums = [2,7,9,3,1]
	// dp[0] = 2, dp[1] = max(2,7) = 7, dp[2] = max(7,2+9) = 11, dp[3] = max(11,7+3) = 11, dp[4] = max(11,11+1) = 12
	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])
	for i := 2; i < n; i++ {
		dp[i] = max(dp[i-1], dp[i-2]+nums[i])
	}

	// return last element
	return dp[n-1]
}

// @lc code=end
