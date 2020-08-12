package ch2linkedlist

import (
    "fmt"
    "math"
)

type ListNode struct {
    Val  int
    Next *ListNode
}

type DNode struct {
    Val  int
    Next *DNode
    Prev *DNode
}

// The 1st Q of ch2: Print same elements in tow linked-list, if is bigger the the other one then move forward, else print
func PrintSameElements(head1 *ListNode, head2 *ListNode) {
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
func RemoveLastKElements(head *ListNode, k int) *ListNode {
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
func RmLastKthNode(head *DNode, k int) *DNode {
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
func RmByMidRatio(head *ListNode, a int, b int) *ListNode {
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
func ReverseSinglyList(head *ListNode) *ListNode {
    var prev *ListNode
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
func ReverseDoubleLinkedList(head *DNode) *DNode {
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

// The 5th Q of ch2: Reverse part linked list
// * 1. 头插法反转链表, 每次将遍历的当前节点插入开始反转的位置。
// * 2. 通过哨兵节点处理 m == 1 的情况
// * 3. 通过m的值确定头插法的头节点位置
// * 4. 通过n-m的值确定执行几次头插操作
func ReversePartLinkedList(head *ListNode, from int, to int) *ListNode {
    dummy, i, j := &ListNode{Next: head}, from, to-from
    d := dummy
    for i > 1 {
        d = d.Next
        i--
    }
    cur := d.Next.Next
    pre := d.Next
    for j > 0 {
        pre.Next = cur.Next
        cur.Next = d.Next
        d.Next = cur
        cur = pre.Next
        j--
    }
    return dummy.Next
}

// The 5th Q of ch2: reverse part lk list
func reverseBetween(head *ListNode, m int, n int) *ListNode {
    res := &ListNode{Next: head}
    node := res
    for i := 1; i < m; i++ {
        node = node.Next
    }

    newHead := node.Next
    var next, prev *ListNode
    for i := m; i <= n; i++ {
        next = newHead.Next
        newHead.Next = prev
        prev = newHead
        newHead = next
    }

    node.Next.Next = next
    node.Next = prev
    return res.Next
}

// The 6th Q of ch2 josephusKill
func josephusKill(head *ListNode, m int) *ListNode {
    if head == nil || head.Next == head || m < 1 {
        return head
    }
    last := head
    for last.Next != head {
        last = last.Next
    }
    cnt := 0
    for head != last {
        cnt++
        if cnt == m {
            last.Next = head.Next
            cnt = 0
        } else {
            last = last.Next
        }
        head = last.Next
    }
    return head
}

type Node struct {
    Val    int
    Next   *Node
    Random *Node
}

// copy random list
func copyRandomList(head *Node) *Node {
    m := map[*Node]*Node{}
    // var m = make(map[*Node]*Node)
    curr := head

    for curr != nil {
        m[curr] = &Node{Val: curr.Val}
        curr = curr.Next
    }
    curr = head
    for curr != nil {
        m[curr].Next = m[curr.Next]
        m[curr].Random = m[curr.Random]
        curr = curr.Next
    }
    return m[head]
}
