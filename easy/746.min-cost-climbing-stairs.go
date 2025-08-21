package code

/*
 * @lc app=leetcode id=746 lang=golang
 *
 * [746] Min Cost Climbing Stairs
 */

// @lc code=start
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minCostClimbingStairs(cost []int) int {
	n := len(cost)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return cost[0]
	}

	// Dynamic programming array
	// dp[i] represents the minimum cost to reach step i
	dp := make([]int, n)

	// Base cases
	dp[0] = cost[0]
	dp[1] = cost[1]

	// Fill the dp array with recurrence formula
	for i := 2; i < n; i++ {
		dp[i] = cost[i] + min(dp[i-1], dp[i-2])
	}

	// Can reach the top from either of the last two steps
	return min(dp[n-1], dp[n-2])
}

// @lc code=end
