/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedArrayToBST(nums []int) *TreeNode {
    return createTree(nums, 0, len(nums) - 1)
}

func createTree(nums []int, start int, end int) *TreeNode {
	if start > end {
		return nil
	}
	mid := (start + end) / 2 
	node := &TreeNode{Val: nums[mid]}
	node.Left = createTree(nums, start, mid-1)
	node.Right = createTree(nums, mid+1, end)
	return node
}