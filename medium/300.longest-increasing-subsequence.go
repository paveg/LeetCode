package code

import (
	"sort"
)

/*
 * @lc app=leetcode id=300 lang=golang
 *
 * [300] Longest Increasing Subsequence
 */

// @lc code=start
func lengthOfLIS(nums []int) int {
	// n := len(nums)

	// if n == 1 {
	// 	return 1
	// }

	// dp := make([]int, n)
	// dp[0] = 1
	// maxLength := 1

	// for i := 1; i < n; i++ {
	// 	dp[i] = 1
	// 	for j := 0; j < i; j++ {
	// 		if nums[i] > nums[j] {
	// 			dp[i] = max(dp[i], dp[j]+1)
	// 		}
	// 	}
	// 	if dp[i] > maxLength {
	// 		maxLength = dp[i]
	// 	}
	// }
	// return maxLength
	n := len(nums)
	if n == 0 {
		return 0
	}

	tails := make([]int, 0, n)
	for _, x := range nums {
		i := sort.SearchInts(tails, x)
		// If i is equal to the length of tails, it means x is larger than any element in tails
		if i == len(tails) {
			tails = append(tails, x)
		} else {
			tails[i] = x
		}
	}

	return len(tails)
}

// @lc code=end
