package code

/*
 * @lc app=leetcode id=70 lang=golang
 *
 * [70] Climbing Stairs
 */

// @lc code=start

// var memo = make(map[int]int)

// top-down approach
// func climbStairsMemo(n int) int {
// 	if n <= 0 {
// 		return 0
// 	}
// 	if n == 1 {
// 		return 1
// 	}
// 	if n == 2 {
// 		return 2
// 	}

// 	if val, ok := memo[n]; ok {
// 		return val
// 	}

// 	result := climbStairs(n-1) + climbStairs(n-2)
// 	memo[n] = result
// 	return result
// }

// bottom-up approach
// O(n) time and O(n) space
func climbStairs(n int) int {
	if n <= 2 {
		return n
	}

	dp := make([]int, n+1)
	dp[1] = 1
	dp[2] = 2

	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[n]
}

// @lc code=end
