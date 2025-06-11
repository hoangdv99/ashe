/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxLevelSum(root *TreeNode) int {
    if root == nil {
        return 0
    }
    max := math.MinInt
    level := 0
    maxLevel := 0
    queue := []*TreeNode{root}

    for len(queue) > 0 {
        sum := 0
        level++
        size := len(queue)
        for range size {
            node := queue[0]
            queue = queue[1:]
            sum += node.Val
            if node.Left != nil {
                queue = append(queue, node.Left)
            }
            if node.Right != nil {
                queue = append(queue, node.Right)
            }
        }
        if sum > max {
            max = sum
            maxLevel = level
        }
    }
    return maxLevel
}