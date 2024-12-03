#
# @lc app=leetcode id=122 lang=python3
#
# [122] Best Time to Buy and Sell Stock II
#

# @lc code=start
class Solution:
    def maxProfit(self, prices: List[int]) -> int:
        profit = 0

        for i, price in enumerate(prices):
            if i > 0 and price > prices[i - 1]:
                profit += price - prices[i - 1]
        return profit
# @lc code=end

