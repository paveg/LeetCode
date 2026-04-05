package code

/*
 * @lc app=leetcode id=200 lang=golang
 *
 * [200] Number of Islands
 */

// @lc code=start
func numIslands(grid [][]byte) int {
	land := byte('1')
	water := byte('0')

	y := len(grid)
	if y == 0 {
		return 0
	}
	x := len(grid[0])
	islandsCount := 0

	var dfs func(i, j int)
	dfs = func(i, j int) {
		if i < 0 || i >= y || j < 0 || j >= x || grid[i][j] == water {
			return
		}
		grid[i][j] = water
		// Explore the four adjacent cells
		dfs(i-1, j)
		dfs(i+1, j)
		dfs(i, j-1)
		dfs(i, j+1)
	}

	for i := 0; i < y; i++ {
		for j := 0; j < x; j++ {
			if grid[i][j] == land {
				islandsCount++
				dfs(i, j)
			}
		}
	}
	return islandsCount
}

// @lc code=end
