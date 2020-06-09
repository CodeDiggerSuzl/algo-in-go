package ch2linkedlist

import (
    "fmt"
    "math"
)

type Node struct {
    Val  int
    Next *Node
}

type DNode struct {
    Val  int
    Next *DNode
    Prev *DNode
}

// The 1st Q of ch2: Print same elements in tow linked-list, if is bigger the the other one then move forward, else print
func printSameElements(head1 *Node, head2 *Node) {
    fmt.Println("Common parts: ")
    if head1.Val > head2.Val {
        head2 = head2.Next
    } else if head1.Val < head2.Val {
        head1 = head1.Next
    } else {
        fmt.Print(head1.Val, "")
        head2 = head2.Next
        head1 = head1.Next
    }
}

// The 2nd Q of ch2: remove the last k element
func removeLastKElements(head *Node, k int) *Node {
    // judge the condition
    if head == nil || k < 1 {
        return head
    }
    cur := head
    for cur != nil {
        k--
        cur = cur.Next
    }

    // len(list) = k
    if k == 0 {
        // remove head
        head = head.Next
    }
    // if k < 0
    if k < 0 {
        cur = head
        for k != 0 {
            k++
            cur = cur.Next
        }
        cur.Next = cur.Next.Next
    }
    return head
}

// The 2ed Q of ch2: remove the last k node of double linked list
func rmLastKthNode(head *DNode, k int) *DNode {
    if k < 1 || head == nil {
        return head
    }
    curr := head
    for curr != nil {
        k--
        curr = curr.Next
    }
    // after a loop
    if k == 0 {
        head = head.Next
        head.Prev = nil
    }
    if k < 0 {
        curr = head
        for k != 0 {
            k++
            curr = curr.Next
        }
        // ! mind the diff
        temp := curr.Next.Next
        curr.Next = temp
        if temp != nil {
            temp.Prev = curr
        }
    }
    return head
}

// The 3rd adv Q of ch2: Delete the node where the pos in a/b
func rmByMidRatio(head *Node, a int, b int) *Node {
    if a < 1 || a > b {
        return head
    }
    // Get l
    cur := head
    l := 0
    for cur != nil {
        l++
        cur = cur.Next
    }
    // 获取要删除的元素
    l = int(math.Ceil(float64(a*l) / float64(b)))

    if l == 1 {
        head = head.Next
    }
    if l > 1 {
        cur = head
        // 不等于 1 就是因为 cur 是要删除元素的前一个元素
        for l != 1 {
            l--
            cur = cur.Next
        }
        cur.Next = cur.Next.Next
    }
    return head
}

// The 4th Q of ch2: reverse singly-list & double-linked-list
func reverseSinglyList(head *Node) *Node {
    var prev *Node
    curr := head
    for curr != nil {
        next := curr.Next
        // reverse
        curr.Next = prev
        // move forward
        prev = curr
        curr = next
    }
    return prev
}

// Reverse the double-linked-list
func reverseDoubleLinkedList(head *DNode) *DNode {
    var prev *DNode
    curr := head
    for curr != nil {
        next := curr.Next
        curr.Next = prev
        curr.Prev = next
        // move forward
        prev = curr
        curr = next
    }
    return prev
}
