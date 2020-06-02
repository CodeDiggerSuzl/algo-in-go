package ch2linkedlist

import "fmt"

type Node struct {
    Val  int
    Next *Node
}

type DNode struct {
    Val  int
    Next *DNode
    Prev *DNode
}

// The 1st Q of ch2:
// Print same elements in tow linked-list
// if is bigger the the other one then move forward, else print
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

// The 2nd Q of ch2:
// remove the last k element
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
func remove(head *DNode, k int) *DNode {
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

