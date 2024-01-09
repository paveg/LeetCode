#
# @lc app=leetcode id=121 lang=python3
#
# [121] Best Time to Buy and Sell Stock
#

# @lc code=start
class Solution:
    # Time Complexity is O(n)
    def maxProfit(self, prices: List[int]) -> int:
        mx = 0
        mn = prices[0]

        for price in prices:
            mn = min(mn, price)
            profit = price - mn
            mx = max(mx, profit)
        return mx
# @lc code=end
