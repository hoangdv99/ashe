/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func searchBST(root *TreeNode, val int) *TreeNode {
    if root == nil {
        return nil
    }
    if root.Val == val {
        return root
    }
    leftSearch := searchBST(root.Left, val)
    if leftSearch != nil {
        return leftSearch
    }
    rightSearch := searchBST(root.Right, val)
    if rightSearch != nil {
        return rightSearch
    }

    return nil
}