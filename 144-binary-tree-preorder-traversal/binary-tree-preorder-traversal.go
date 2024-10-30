/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func preorderTraversal(root *TreeNode) []int {
    result := []int{}
    if root == nil {
        return result
    }
    traverse(&result, root)
    return result
}

func traverse(list *[]int, root *TreeNode) {
    if root != nil {
        *list = append(*list, root.Val)
    }
    if root.Left != nil {
        traverse(list, root.Left)
    }
    if root.Right != nil {
        traverse(list, root.Right)
    }
}