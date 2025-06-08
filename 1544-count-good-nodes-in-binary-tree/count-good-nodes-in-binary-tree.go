/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func goodNodes(root *TreeNode) int {
    return dfs(root, math.MinInt)
}

func dfs(root *TreeNode, max int) int {
    if root == nil {
        return 0
    }

    count := 0
    if root.Val >= max {
        count++
        max = root.Val
    }

    return count + dfs(root.Left, max) + dfs(root.Right, max)
}