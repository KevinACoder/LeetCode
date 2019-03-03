/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {

}

func reverseNextK(head *ListNode, k int) *ListNode {
	curr, tail := head.Next, head.Next
	for 
}

func reverseList(head, tail *ListNode) *ListNode {
	curr, prev := head, tail
	for curr != tail {
		next := cur.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}