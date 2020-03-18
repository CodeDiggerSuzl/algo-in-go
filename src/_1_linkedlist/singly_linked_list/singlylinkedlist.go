package singly_linked_list

import "fmt"

// One-way linked list

// node
type ListNode struct {
	next  *ListNode
	value interface{}
}

// singly linked list
type LinkedList struct {
	head   *ListNode
	length uint
}

// new List Node
func NewListNode(v interface{}) *ListNode {
	return &ListNode{nil, v}
}

// get next node
func (a *ListNode) GetNext() *ListNode {
	return a.next
}

// get node value
func (a *ListNode) GetValue() interface{} {
	return a.next
}

// methods of linked-list
func NewLinkedList() *LinkedList {
	return &LinkedList{NewListNode(0), 0}
}

// insert node after a node
// v-> the insert value
func (a *LinkedList) InsertAfter(p *ListNode, v interface{}) bool {
	if nil == p {
		return false
	}
	// create a node
	newNode := NewListNode(v)
	oldNext := p.next
	p.next = newNode
	// the new node's next now is the old next
	newNode.next = oldNext
	a.length++
	return true
}

// insert b4 one node
func (a *LinkedList) InsertBefore(p *ListNode, v interface{}) bool {
	// judge
	if p == nil || p == a.head {
		return false
	}
	// need to check if it's the
	cur := a.head.next
	pre := a.head

	// if origin linked list only has one node
	// if head's next is not nil
	for nil != cur {
		if cur == p {
			break
		}
		pre = cur
		cur = cur.next
	}
	if nil == cur {
		return false
	}

}

func (a *LinkedList) InsertToHead(p *ListNode, v interface{}) bool {
	return a.InsertAfter(a.head, v)
}

// delete the node
func (a *LinkedList) DeleteNode(p *ListNode) bool {
	if p == nil {
		return false
	}
	curNext := a.head.next
	preHead := a.head

	for curNext != nil {
		if curNext == p {
			break
		}
		preHead = curNext
		curNext = curNext.next
	}

	if curNext == nil {
		return false
	}
	preHead.next = p.next
	p = nil
	a.length--
	return true
}

// print list
func (a *LinkedList) Print() {
	cur := a.head.next
	p := ""
	for nil != cur {
		p += fmt.Sprintf("%+v", cur.GetValue())
		cur = cur.next
		if nil != cur {
			p += "->"
		}
	}
	fmt.Println(p)
}
