#
# @lc app=leetcode id=144 lang=python3
#
# [144] Binary Tree Preorder Traversal
#

# @lc code=start
# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def preorderTraversal(self, root: Optional[TreeNode]) -> List[int]:
        def preorder(node: Optional[TreeNode]):
            if node is None:
                return
            ans.append(node.val)
            preorder(node.left)
            preorder(node.right)

        ans = []
        preorder(root)
        return ans

# @lc code=end
