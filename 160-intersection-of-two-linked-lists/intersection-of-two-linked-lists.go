/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
    if headA == nil || headB == nil {
        return nil
    }

    m := make(map[*ListNode]int)

    for headA != nil {
        m[headA] = 1
        headA = headA.Next
    }

    for headB != nil {
        _, ok := m[headB]
        if ok {
            return headB
        }
        headB = headB.Next
    }
    return nil
}