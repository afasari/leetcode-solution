# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def isSameTree(self, p: Optional[TreeNode], q: Optional[TreeNode]) -> bool:
        if p is None or q is None:
            return p == q
            
        # if not p and not q:
        #     print("and")
        #     return True

        # if not p or not q or p.val != q.val:
        #     return False
            

        return p.val == q.val and self.isSameTree(p.left, q.left) and self.isSameTree(p.right, q.right)
        
        # stack = [(p, q)]
        # while stack:
        #     p, q = stack.pop()
        #     if p or q:
        #         if not p or not q or p.val != q.val:
        #             return False
        #         stack.append((p.left, q.left))
        #         stack.append((p.right, q.right))
        # return True