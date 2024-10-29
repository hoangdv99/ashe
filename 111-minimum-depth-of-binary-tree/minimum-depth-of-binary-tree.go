/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func minDepth(root *TreeNode) int {
    if root == nil {
        return 0
    }
    if root.Left == nil && root.Right == nil {
        return 1
    }

    leftMinDepth := 1 + minDepth(root.Left)
    rightMinDepth := 1 + minDepth(root.Right)
    if root.Left == nil {
        return rightMinDepth
    } else if root.Right == nil {
        return leftMinDepth
    } else {
        if rightMinDepth < leftMinDepth {
            return rightMinDepth
        }
        return leftMinDepth
    }
}