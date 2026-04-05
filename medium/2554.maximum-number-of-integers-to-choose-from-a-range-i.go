package code

/*
 * @lc app=leetcode id=2554 lang=golang
 *
 * [2554] Maximum Number of Integers to Choose From a Range I
 */

// @lc code=start
func maxCount(banned []int, n int, maxSum int) int {
	mc := 0

	bannedSet := make(map[int]struct{})
	for _, b := range banned {
		bannedSet[b] = struct{}{}
	}

	sum := 0
	// [1, 2, 3, ..., n]
	for i := 1; i <= n; i++ {
		_, banned := bannedSet[i]
		if !banned && sum+i <= maxSum {
			sum += i
			mc++
		}
	}
	return mc
}

// @lc code=end
