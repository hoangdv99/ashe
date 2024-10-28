/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isBalanced(root *TreeNode) bool {
    _, balanced := checkHeight(root)
    return balanced
}

func checkHeight(root *TreeNode) (float64, bool) {
    if root == nil {
        return 0, true
    }
    leftHeight, leftBalanced := checkHeight(root.Left)
    rightHeight, rightBalanced := checkHeight(root.Right)
    currentBalanced := leftBalanced && rightBalanced && math.Abs(leftHeight - rightHeight) <= 1

    return math.Max(leftHeight, rightHeight) + 1, currentBalanced
}