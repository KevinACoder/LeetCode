/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
	dummy := ListNode{Val: 0, Next: head}
	p := &dummy
	for p != nil {
		p = reverseNextK(p, k)
	}
	return dummy.Next
}

func reverseNextK(head *ListNode, k int) *ListNode {
	curr, tail := head.Next, head.Next
	for i := 0; i < k; i++ {
		if tail == nil {
			return nil
		}
		tail = tail.Next
	}
	head.Next = reverseList(curr, tail)
	return curr
}

func reverseList(head, tail *ListNode) *ListNode {
	curr, prev := head, tail
	for curr != tail {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}