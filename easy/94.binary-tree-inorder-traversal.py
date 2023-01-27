#
# @lc app=leetcode id=94 lang=python3
#
# [94] Binary Tree Inorder Traversal
#

# @lc code=start
# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def inorderTraversal(self, root: Optional[TreeNode]) -> List[int]:
        def inorder(node: Optional[TreeNode]):
            if node is None:
                return
            inorder(node.left)
            ans.append(node.val)
            inorder(node.right)

        ans = []
        inorder(root)
        return ans

# @lc code=end
