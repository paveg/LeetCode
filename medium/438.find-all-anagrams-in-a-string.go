package code

/*
 * @lc app=leetcode id=438 lang=golang
 *
 * [438] Find All Anagrams in a String
 */

// @lc code=start
func findAnagrams(s string, p string) []int {
	if len(s) < len(p) {
		return nil
	}

	result := []int{}
	var sCount, pCount [26]int
	for i := 0; i < len(p); i++ {
		pCount[p[i]-'a']++
		sCount[s[i]-'a']++
	}
	if sCount == pCount {
		result = append(result, 0)
	}

	for i := len(p); i < len(s); i++ {
		sCount[s[i]-'a']++
		sCount[s[i-len(p)]-'a']--
		if sCount == pCount {
			result = append(result, i-len(p)+1)
		}
	}
	return result
}

// @lc code=end
