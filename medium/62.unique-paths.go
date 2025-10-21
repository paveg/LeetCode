package code

/*
 * @lc app=leetcode id=62 lang=golang
 *
 * [62] Unique Paths
 */

// @lc code=start

func uniquePaths(m int, n int) int {
	rows := m
	if rows == 0 {
		return 0
	}
	cols := n

	dp := make([]int, cols)

	for r := 0; r < m; r++ {
		for c := 0; c < n; c++ {
		}
	}
	return dp[(m-1)*n+(n-1)]
}

// @lc code=end
