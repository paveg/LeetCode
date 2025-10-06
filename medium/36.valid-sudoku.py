#
# @lc app=leetcode id=36 lang=python3
#
# [36] Valid Sudoku
#

# @lc code=start
class Solution:
    # TODO: You need to understand the problem and this solution.
    def isValidSudoku(self, board: List[List[str]]) -> bool:
        m = {}
        for i in range(9):
            for j in range(9):
                if board[i][j] != '.':
                    num = board[i][j]
                    if num not in m:
                        m[num] = []
                    m[num].append((i, j))
        for num in m:
            if len(m[num]) > 1:
                for i in range(len(m[num])):
                    print(m[num][i])
                    for j in range(i + 1, len(m[num])):
                        if m[num][i][0] == m[num][j][0] or m[num][i][1] == m[num][j][1] or (m[num][i][0] // 3 == m[num][j][0] // 3 and m[num][i][1] // 3 == m[num][j][1] // 3):
                            return False
        return True   
# @lc code=end

