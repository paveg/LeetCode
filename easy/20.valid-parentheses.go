package code

/*
 * @lc app=leetcode id=20 lang=golang
 *
 * [20] Valid Parentheses
 */

// @lc code=start
func isValid(s string) bool {
	if len(s)%2 != 0 {
		return false
	}

	pair := map[rune]rune{')': '(', '}': '{', ']': '['}
	stack := []rune{}

	for _, char := range s {
		if open, exists := pair[char]; exists {
			if len(stack) == 0 || stack[len(stack)-1] != open {
				return false
			}
			// Pop the last element
			stack = stack[:len(stack)-1]
		} else {
			// Push the current element
			stack = append(stack, char)
		}
	}

	return len(stack) == 0
}

// @lc code=end
