/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func listToSlice(l *ListNode) (ints []int) {
    for {
        ints = append(ints, l.Val)
        if l.Next == nil {
            break
        }
        l = l.Next
    }
    
    return
}

func addSlices(s1, s2 []int) (out []int) {
    if len(s2) > len(s1) {
        s1, s2 = s2, s1
    }
    
    remainder := 0
    for i, v1 := range s1 {
        v2 := 0
        if i < len(s2) {
            v2 = s2[i]
        }
        total := (v1 + v2 + remainder)
        out = append(out, total % 10)
        remainder = total / 10 
    }
    if remainder > 0 {
        out = append(out, remainder)
    }
    
    
    return
}

func sliceToList(in []int) *ListNode {
    var node *ListNode
    for i := len(in) - 1; i >= 0; i-- {
        node = &ListNode{in[i], node}
    }
    
    return node
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    // return sliceToList(addSlices(listToSlice(l1), listToSlice(l2)))
    
    return solution(l1, l2)
}


func solution(l1 *ListNode, l2 *ListNode) *ListNode {
	dummyNode := ListNode{0, nil}
	carry := 0
	current := &dummyNode

	for l1 != nil || l2 != nil {
		var x, y int
		if (l1 == nil) {
			x = 0
		} else {
			x = l1.Val;
		}

		if (l2 == nil) {
			y = 0
		} else {
			y = l2.Val
		}

		sum := x + y + carry

		current.Next = &ListNode{sum % 10, nil}
		carry = sum / 10

		current = current.Next
		if l1 != nil {
			l1 = l1.Next
		}

		if l2 != nil {
			l2 = l2.Next
		}
	}

	if carry != 0 {
		current.Next = &ListNode{1, nil}
	}

	return dummyNode.Next
}