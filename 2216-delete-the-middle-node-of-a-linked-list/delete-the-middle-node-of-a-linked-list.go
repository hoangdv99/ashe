/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteMiddle(head *ListNode) *ListNode {
    if head.Next == nil {
        return nil
    }

    slow, fast := head, head
    prevSlow := slow
    for fast.Next != nil {
        if fast.Next.Next == nil {
            fast = fast.Next
        } else {
            fast = fast.Next.Next
        }
        prevSlow = slow
        slow = slow.Next
        if fast.Next == nil {
            prevSlow.Next = slow.Next
        }
    }
   
    return head
}