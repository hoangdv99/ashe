/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func pairSum(head *ListNode) int {
    fast, slow := head, head
    for fast != nil && fast.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
    }

    cur := slow
    var prev *ListNode
    for cur != nil {
        next := cur.Next
        cur.Next = prev
        prev = cur
        cur = next
    }

    max := 0
    first, second := head, prev
    for second != nil {
        sum := first.Val + second.Val
        if sum > max {
            max = sum
        }
        first = first.Next
        second = second.Next
    }

    return max
}