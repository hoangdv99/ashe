/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum(root *TreeNode, targetSum int) int {
    total := 0
    
    var helper func(node *TreeNode, cur int)

    helper = func(node *TreeNode, cur int) {
        if node == nil {
            return
        }
        if cur + node.Val == targetSum {
            total++
        }

        helper(node.Left, cur + node.Val)
        helper(node.Right, cur + node.Val)
    }

    var dfs func(node *TreeNode)
    dfs = func(node *TreeNode) {
        if node == nil {
            return
        }
        helper(node, 0)
        dfs(node.Left)
        dfs(node.Right)
    }

    dfs(root)

    return total
}