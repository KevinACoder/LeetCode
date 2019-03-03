/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	return mergeKLists(lists, 0, len(lists)-1)
}

func mergeKLists(lists []*ListNode, start, end int) *ListNode {
	if start == end {
		return lists[start]
	}

	mid := start + (end-start)/2
	left := mergeKLists(lists, start, mid)
	right := mergeKLists(lists, mid+1, end)
	return mergeLists(left, right)
}

func mergeLists(l1, l2 *ListNode) *ListNode {
	dummy := ListNode{Val: 0, Next: nil}
	p := &dummy
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			p.Next = l1
			l1 = l1.Next
		} else {
			p.Next = l2
			l2 = l2.Next
		}
		p = p.Next
	}

	if l1 != nil {
		p.Next = l1
	} else if l2 != nil {
		p.Next = l2
	}

	return dummy.Next
}