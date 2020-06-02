package linkedlist

/**
 *  删除链表中等于给定值 val 的所有节点。
 *
 *  示例:
 *
 *  输入: 1->2->6->3->4->5->6, val = 6
 *  输出: 1->2->3->4->5
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

func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return head
	}
	// 似乎不能这样处置
	watch := &ListNode{0, head}
	cur := head
	for cur != nil {
		if cur.Val == val {
			watch.Next = cur.Next
		} else {
			watch = cur
			cur = cur.Next
		}
	}
	return watch.Next
}

func removeElements2(head *ListNode, val int) *ListNode {
	watch := new(ListNode)
	watch.Next = head
	pre, cur := watch, head
	for cur != nil {
		if cur.Val == val {
			pre.Next = cur.Next
			// 去掉下面这一行就不行了
			cur = cur.Next
		} else {
			pre = cur
			cur = cur.Next
		}
	}
	return watch.Next
}
