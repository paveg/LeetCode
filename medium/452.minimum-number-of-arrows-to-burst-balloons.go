package code

import "sort"

/*
 * @lc app=leetcode id=452 lang=golang
 *
 * [452] Minimum Number of Arrows to Burst Balloons
 */

// @lc code=start
func findMinArrowShots(points [][]int) int {
	// [10,16],[2,8],[1,6],[7,12]
	// 0 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16
	//                     ====================
	// 	   =============
	//   ===========
	//               ==============
	result := 0
	sort.Slice(points, func(i, j int) bool {
		return points[i][1] < points[j][1]
	})
	result++
	xEnd := points[0][1]
	for i := 1; i < len(points); i++ {
		if points[i][0] > xEnd {
			result++
			xEnd = points[i][1]
		}
	}

	return result
}

// @lc code=end
