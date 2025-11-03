package code

/*
 * @lc app=leetcode id=102 lang=golang
 *
 * [102] Binary Tree Level Order Traversal
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
	// If the root is nil, return an empty slice
	if root == nil {
		return [][]int{}
	}

	result := [][]int{}
	queue := []*TreeNode{root}
	// Breadth-First Search (BFS)
	for len(queue) > 0 {
		level := []int{}
		// Process the current nodes in the queue
		for _, node := range queue {
			level = append(level, node.Val)
		}
		result = append(result, level)

		nextQueue := []*TreeNode{}
		// Enqueue the children of the current nodes
		for _, node := range queue {
			if node.Left != nil {
				nextQueue = append(nextQueue, node.Left)
			}
			if node.Right != nil {
				nextQueue = append(nextQueue, node.Right)
			}
		}
		// Move to the next level
		queue = nextQueue
	}

	// Return the final result
	return result
}

// @lc code=end
