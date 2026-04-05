//go:build local
// +build local

package code

// ローカル専用の ListNode 定義（LeetCode 側では読み込まれない）
type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
