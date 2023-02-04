#
# @lc app=leetcode id=404 lang=python3
#
# [404] Sum of Left Leaves
#

# @lc code=start
# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def sumOfLeftLeaves(self, root: Optional[TreeNode]) -> int:
        self.ans = 0
        def dig(node: Optional[TreeNode], x: str):
            if (node is None):
                return
            if (not node.left and not node.right):
                if (x == 'l'):
                    self.ans += node.val
            if node.left:
                dig(node.left, 'l')
            if node.right:
                dig(node.right, 'r')

        dig(root, '')
        return self.ans
# @lc code=end
