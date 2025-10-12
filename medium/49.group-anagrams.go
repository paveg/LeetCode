package code

import (
	"fmt"
	"sort"
)

/*
 * @lc app=leetcode id=49 lang=golang
 *
 * [49] Group Anagrams
 */

// @lc code=start
func groupAnagrams(strs []string) [][]string {
	m := map[string]int{}
	anagrams := [][]string{}

	i := 0

	for _, s := range strs {
		sorted := strSort(s)
		v, ok := m[sorted]
		if ok {
			// exists
			anagrams[v] = append(anagrams[v], s)
		} else {
			anagrams = append(anagrams, []string{s})
			m[sorted] = i
			i++
		}
	}

	return anagrams
}

func strSort(s string) string {
	b := []byte(s)
	sort.Slice(b, func(i, j int) bool { return b[i] < b[j] })
	return string(b)
}

// @lc code=end
