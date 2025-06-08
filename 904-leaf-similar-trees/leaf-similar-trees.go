/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
    leaves1 := []int{}
    leaves2 := []int{}
    getLeaves(root1, &leaves1)
    getLeaves(root2, &leaves2)

    fmt.Println(leaves1)

    if len(leaves1) != len(leaves2) {
        return false
    }

    for i := range leaves1 {
        if leaves1[i] != leaves2[i] {
            return false
        }
    }
    return true
}

func getLeaves(root *TreeNode, leaves *[]int) {
    if root == nil {
        return
    }
    if root.Left == nil && root.Right == nil {
        *leaves = append(*leaves, root.Val)
    }

    getLeaves(root.Left, leaves)
    getLeaves(root.Right, leaves)
}