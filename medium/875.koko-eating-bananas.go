package code

/*
 * @lc app=leetcode id=875 lang=golang
 *
 * [875] Koko Eating Bananas
 */

// @lc code=start
func minEatingSpeed(piles []int, h int) int {
	maxPile := 0
	for _, pile := range piles {
		if pile > maxPile {
			maxPile = pile
		}
	}

	left, right := 1, maxPile
	result := right

	for left <= right {
		mid := left + (right-left)/2
		totalHours := 0

		for _, pile := range piles {
			hours := pile / mid
			if pile%mid != 0 {
				hours++
			}
			totalHours += hours
		}

		if totalHours <= h {
			if mid < result {
				result = mid
			}
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return result
}

// @lc code=end
