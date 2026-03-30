package code

/*
 * @lc app=leetcode id=1529 lang=golang
 *
 * [1529] Minimum Suffix Flips
 */

// @lc code=start
func minFlips(target string) int {
	flips := 0
	current := byte('0')

	for i := 0; i < len(target); i++ {
		if target[i] != current {
			flips++
			current = target[i]
		}
	}
	return flips
}

// @lc code=end
