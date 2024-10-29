/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	if root.Val == targetSum && root.Left == nil && root.Right == nil {
		return true
	}
	hasLeftPathSum := hasPathSum(root.Left, targetSum-root.Val)
	hasRightPathSum := hasPathSum(root.Right, targetSum-root.Val)
	return hasLeftPathSum || hasRightPathSum
}