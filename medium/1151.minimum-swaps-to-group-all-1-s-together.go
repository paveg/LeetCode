package code

/*
 * @lc app=leetcode id=1151 lang=golang
 *
 * [1151] Minimum Swaps to Group All 1's Together
 */

// @lc code=start
func minSwaps(data []int) int {
	// Count the number of 1's in the array
	ones := 0
	for _, v := range data {
		ones += v
	}
	// If there are no 1's, no swaps are needed, so return 0
	// If there is only one 1, no swaps are needed, so return 0
	if ones == 0 || ones == 1 {
		return 0
	}
	zeros := 0
	for i := 0; i < ones; i++ {
		if data[i] == 0 {
			zeros++
		}
	}
	minZeros := zeros

	for i := ones; i < len(data); i++ {
		if data[i] == 0 {
			zeros++
		}
		if data[i-ones] == 0 {
			zeros--
		}
		if zeros < minZeros {
			minZeros = zeros
		}
	}
	return minZeros
}

// @lc code=end
