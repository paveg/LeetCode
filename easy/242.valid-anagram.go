package code

/*
 * @lc app=leetcode id=242 lang=golang
 *
 * [242] Valid Anagram
 */

// @lc code=start
func isAnagram(s string, t string) bool {
	count := [26]int{}
	for i := 0; i < len(s); i++ {
		count[s[i]-'a']++
	}

	for i := 0; i < len(t); i++ {
		count[t[i]-'a']--
	}

	for i := 0; i < 26; i++ {
		if count[i] != 0 {
			return false
		}
	}
	return true
}

// @lc code=end
