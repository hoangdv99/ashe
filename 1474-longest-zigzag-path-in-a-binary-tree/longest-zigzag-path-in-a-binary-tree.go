/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func longestZigZag(root *TreeNode) int {
    if root == nil || (root.Left == nil && root.Right == nil) {
        return 0
    }
    maxLength := 0

    var dfs func(node *TreeNode, isLeft bool, length int)
    dfs = func(node *TreeNode, isLeft bool, length int) {
        if node == nil {
            return
        }
        if length > maxLength {
            maxLength = length
        }
        if isLeft {
            dfs(node.Left, false, length+1)
            dfs(node.Right, true, 1)
        } else {
            dfs(node.Right, true, length+1)
            dfs(node.Left, false, 1)
        }
    }
    dfs(root.Left, false, 1)
    dfs(root.Right, true, 1)

    return maxLength
}