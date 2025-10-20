package code

/*
 * @lc app=leetcode id=3 lang=golang
 *
 * [3] Longest Substring Without Repeating Characters
 */

// @lc code=start
func lengthOfLongestSubstring(s string) int {
	left, best := 0, 0
	charIndex := make(map[rune]int)

	// Time complexity: O(n)
	// Space complexity: O(min(m,n)) where m is the size of the charset, n is the size of the string
	for right, char := range s {
		if idx, found := charIndex[char]; found && idx >= left {
			left = idx + 1
		}
		charIndex[char] = right

		if cur := right - left + 1; cur > best {
			best = cur
		}
	}

	return best
}

// @lc code=end
