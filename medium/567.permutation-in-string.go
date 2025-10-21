package code

/*
 * @lc app=leetcode id=567 lang=golang
 *
 * [567] Permutation in String
 */

// @lc code=start
func checkInclusion(s1 string, s2 string) bool {
	// Time complexity: O(n)
	// Space complexity: O(1)
	left := 0
	count := [26]int{}

	if len(s1) > len(s2) {
		return false
	}

	for i := 0; i < len(s1); i++ {
		count[s1[i]-'a']++
	}

	for right := 0; right < len(s2); right++ {
		count[s2[right]-'a']--
		for count[s2[right]-'a'] < 0 {
			count[s2[left]-'a']++
			left++
		}

		if right-left+1 == len(s1) {
			return true
		}
	}

	return false
}

// @lc code=end
