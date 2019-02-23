/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := ListNode{Val: 0, Next: nil}
	p := &dummy
	var carry int
	for l1 != nil || l2 != nil || carry > 0 {
		var sum int
		sum += carry
		if l1 != nil{
			sum += l1->Val
			l1 = l1->Next
		}

		if l2 != nil{
			sum += l2->Val
			l2 = l2->Next
		}

		carry = sum/10
		sum = sum%10
		p->Next = new ListNode{Val: sum, Next: nil}
		p = p->Next
	}

	return dummy.Next
}