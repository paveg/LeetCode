#
# @lc app=leetcode id=119 lang=python3
#
# [119] Pascal's Triangle II
#

# @lc code=start
class Solution:
    def getRow(self, rowIndex: int) -> List[int]:
        def pascalsTriangle(rowNums: int) -> List[List[int]]:
            ans = []
            for i in range(rowNums):
                ans.append([])
                ans[i].append(1)
                for j in range(1, i):
                    ans[i].append(ans[i-1][j-1] + ans[i-1][j])
                if i != 0:
                    ans[i].append(1)
            return ans
        result = pascalsTriangle(rowIndex+1)
        return result[rowIndex]
# @lc code=end
