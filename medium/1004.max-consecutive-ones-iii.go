package code

/*
 * @lc app=leetcode id=1004 lang=golang
 *
 * [1004] Max Consecutive Ones III
 */

// @lc code=start
func longestOnes(nums []int, k int) int {
	// left, best := 0, 0
	// count := [2]int{}

	// for right := 0; right < len(nums); right++ {
	// 	if nums[right] == 0 {
	// 		count[0]++
	// 	} else if nums[right] == 1 {
	// 		count[1]++
	// 	}

	// 	Current window size is right - left + 1
	// 	Number of zeros to change is count[0]
	// 	if count[0] > k {
	// 		if nums[left] == 0 {
	// 			count[0]--
	// 		} else if nums[left] == 1 {
	// 			count[1]--
	// 		}
	// 		left++
	// 	}

	// 	if cur := right - left + 1; cur > best {
	// 		best = cur
	// 	}
	// }

	// return best
	// Improved version
	left, best, zeroCount := 0, 0, 0

	// Time complexity: O(n)
	// Space complexity: O(1)
	for right := 0; right < len(nums); right++ {
		if nums[right] == 0 {
			zeroCount++
		}

		if zeroCount > k {
			if nums[left] == 0 {
				zeroCount--
			}
			left++
		}

		if cur := right - left + 1; cur > best {
			best = cur
		}
	}

	return best
}

// @lc code=end
