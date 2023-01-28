#
# @lc app=leetcode id=118 lang=python3
#
# [118] Pascal's Triangle
#

# @lc code=start
class Solution:
    def generate(self, numRows: int) -> List[List[int]]:
        ans = []
        for i in range(numRows):
            ans.append([])
            ans[i].append(1)
            for j in range(1, i):
                ans[i].append(ans[i-1][j-1] + ans[i-1][j])
            if i != 0:
                ans[i].append(1)
        return (ans)
# @lc code=end
