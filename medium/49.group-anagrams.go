package code

import (
	"sort"
)

/*
 * @lc app=leetcode id=49 lang=golang
 *
 * [49] Group Anagrams
 */

// @lc code=start
func groupAnagrams(strs []string) [][]string {
	// O(n * k log k)
	groups := make(map[string][]string, len(strs))
	for _, s := range strs {
		k := sortStringASCII(s)
		// groups["aet"] = ["eat", "tea", "ate"]
		groups[k] = append(groups[k], s)
	}

	// map to 2d slice
	res := make([][]string, 0, len(groups))
	for _, g := range groups {
		res = append(res, g)
	}

	return res
}

// func groupAnagramsWithfrequencyArrayKey(strs []string) [][]string {
// 	// O(n * k)
// 	groups := make(map[[26]byte][]string, len(strs))
// 	for _, s := range strs {
// 		var key [26]byte
// 		for i := 0; i < len(s); i++ {
// 			//
// 			key[s[i]-'a']++
// 		}
// 		groups[key] = append(groups[key], s)
// 	}

// 	res := make([][]string, 0, len(groups))
// 	for _, g := range groups {
// 		res = append(res, g)
// 	}

// 	return res
// }

func sortStringASCII(s string) string {
	b := []byte(s)
	sort.Slice(b, func(i, j int) bool {
		return b[i] < b[j]
	})
	return string(b)
}

// @lc code=end
