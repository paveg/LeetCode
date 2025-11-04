package code

import "sort"

/*
 * @lc app=leetcode id=56 lang=golang
 *
 * [56] Merge Intervals
 */

// @lc code=start
func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return intervals
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	startIdx, endIdx := 0, 1

	// Merge overlapping intervals
	// [1,3]
	merged := make([][]int, 0, len(intervals))
	merged = append(merged, intervals[0])
	for i := 1; i < len(intervals); i++ {
		// 2 <= 3
		if intervals[i][startIdx] <= merged[len(merged)-1][endIdx] {
			// Overlapping intervals, merge them
			// [1, max(3,6)]
			merged[len(merged)-1][endIdx] = max(merged[len(merged)-1][endIdx], intervals[i][endIdx])
		} else {
			// Non-overlapping interval, add it to the result
			merged = append(merged, intervals[i])
		}
	}

	return merged
}

// @lc code=end
