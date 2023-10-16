/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    // if l1 == nil {
	// 	return l2
	// }
	// if l2 == nil {
	// 	return l1
	// }
	// if l1.Val < l2.Val {
	// 	l1.Next = mergeTwoLists(l1.Next, l2)
	// 	return l1
	// }
	// l2.Next = mergeTwoLists(l1, l2.Next)
	// return l2

	res := &ListNode{}
	cur := res
	for list1 != nil || list2 != nil {
		if list1 == nil {
			cur.Next = list2
			list2 = nil
			continue
		}
		if list2 == nil {
			cur.Next = list1
			list1 = nil
			continue
		}

		if list1.Val > list2.Val {
			cur.Next = list2
			cur, list2 = cur.Next, list2.Next
		} else {
			cur.Next = list1
			cur, list1 = cur.Next, list1.Next
		}
	}
	return res.Next
}