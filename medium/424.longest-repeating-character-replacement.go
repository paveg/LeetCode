package code

/*
 * @lc app=leetcode id=424 lang=golang
 *
 * [424] Longest Repeating Character Replacement
 */

// @lc code=start
func characterReplacement(s string, k int) int {
	// Time complexity: O(n)
	// Space complexity: O(1)
	count := [26]int{}

	// left: window left index
	// right: window right index
	// best: longest valid window size
	// maxFreq: max frequency of a single character in the current window
	// (right-left+1): current window size
	// (right-left+1)-maxFreq: number of letters to change to make all characters in the window the same
	left, best, maxFreq := 0, 0, 0

	for right := 0; right < len(s); right++ {
		idx := s[right] - 'A'
		count[idx]++

		if count[idx] > maxFreq {
			maxFreq = count[idx]
		}

		// Current window size is right - left + 1
		// Number of letters to change is (right - left + 1) - max
		if (right-left+1)-maxFreq > k {
			count[s[left]-'A']--
			left++
		}

		if cur := right - left + 1; cur > best {
			best = cur
		}
	}

	return best
}

// @lc code=end
