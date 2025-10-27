package code

/*
 * @lc app=leetcode id=121 lang=golang
 *
 * [121] Best Time to Buy and Sell Stock
 */

// @lc code=start
func maxProfit(prices []int) int {
	max := 0
	min := prices[0]

	profit := 0

	for _, price := range prices {
		if price < min {
			min = price
		}
		profit = price - min
		if profit > max {
			max = profit
		}
	}
	return max
}

// @lc code=end
