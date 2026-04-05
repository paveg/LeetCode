//go:build ignore

package code

import "fmt"

/*
 * @lc app=leetcode id=2134 lang=golang
 *
 * [2134] Minimum Swaps to Group All 1's Together II
 */

// @lc code=start
func minSwaps(nums []int) int {
	ones := 0
	for _, v := range nums {
		ones += v
	}
	fmt.Println(ones)
	if ones == 0 {
		return 0
	}

	n := len(nums)
	zeros := 0
	for i := 0; i < ones; i++ {
		if nums[i] == 0 {
			zeros++
		}
	}
	minZeros := zeros

	for i := 0; i < n; i++ {
		if nums[(i+ones)%n] == 0 {
			zeros++
		}
		if nums[i] == 0 {
			zeros--
		}
		if zeros < minZeros {
			minZeros = zeros
		}
	}
	return minZeros
}

// @lc code=end
