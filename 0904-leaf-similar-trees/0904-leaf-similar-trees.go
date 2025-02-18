/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
    l1 := make([]int, 0)
    l2 := make([]int, 0)

    dp(root1, &l1)
    dp(root2, &l2)
    return reflect.DeepEqual(l1, l2)
}

func dp(node *TreeNode, leafList *[]int) {
    if node == nil{
        return
    }

    if node.Left == nil && node.Right == nil{
        *leafList = append(*leafList, node.Val)
        return
    }

    dp(node.Left, leafList)
    dp(node.Right, leafList)
}

