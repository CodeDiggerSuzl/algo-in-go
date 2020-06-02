package linkedlist

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
    Val  int
    Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }
    next := head.Next
    head.Next = nil
    newHead := reverseList(next)
    next.Next = head
    return newHead
}

// Reverse a list with loop
func reverseListWithLoop(head *ListNode) *ListNode {
    var prev *ListNode
    curr := head
    for curr.Next != nil {
        temp := curr.Next
        curr.Next = prev
        prev = curr
        curr = temp
    }
    return prev
}
