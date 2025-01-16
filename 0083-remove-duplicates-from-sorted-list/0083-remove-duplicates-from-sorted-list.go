/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
    if head == nil{
        return head
    }

    current := head

    for current != nil{
        if current.Next != nil && current.Val == current.Next.Val{
            current.Next = current.Next.Next
        } else {
            current = current.Next
        }
    }

    return head
    
}