package code

/*
 * @lc app=leetcode id=64 lang=golang
 *
 * [64] Minimum Path Sum
 */

// @lc code=start
func minPathSum(grid [][]int) int {
	// m, n := len(grid), len(grid[0])
	// // grid = [[1,3,1],[1,5,1],[4,2,1]]
	// // [[1,4,5],[2,7,8],[6,8,9]]
	// for i := 0; i < m; i++ {
	// 	for j := 0; j < n; j++ {
	// 		if i == 0 && j == 0 {
	// 			continue
	// 		}
	// 		if i == 0 {
	// 			grid[i][j] += grid[i][j-1]
	// 		} else if j == 0 {
	// 			grid[i][j] += grid[i-1][j]
	// 		} else {
	// 			grid[i][j] += min(grid[i-1][j], grid[i][j-1])
	// 		}
	// 	}
	// }
	// return grid[m-1][n-1]

	// Optimized space complexity O(n) using a 1D dp array
	m, n := len(grid), len(grid[0])
	dp := make([]int, n)

	dp[0] = grid[0][0]
	// Initialize the first row
	for j := 1; j < n; j++ {
		dp[j] = dp[j-1] + grid[0][j]
	}

	// Fill the dp array
	for i := 1; i < m; i++ {
		// Update the first column for the current row
		dp[0] += grid[i][0]
		for j := 1; j < n; j++ {
			dp[j] = min(dp[j], dp[j-1]) + grid[i][j]
		}
	}

	return dp[n-1]
}

// @lc code=end
