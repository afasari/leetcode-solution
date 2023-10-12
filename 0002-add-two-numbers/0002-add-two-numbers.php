/**
 * Definition for a singly-linked list.
 * class ListNode {
 *     public $val = 0;
 *     public $next = null;
 *     function __construct($val = 0, $next = null) {
 *         $this->val = $val;
 *         $this->next = $next;
 *     }
 * }
 */
class Solution {

    /**
     * @param ListNode $l1
     * @param ListNode $l2
     * @return ListNode
     */
    function addTwoNumbers($l1, $l2) {
        $root = new ListNode();
        $extra = 0;

        do {
            $sum = (int) $l1?->val + (int) $l2?->val + $extra;
            $extra = floor ($sum / 10);

            if (empty($current)) {
                $current = $root;
                $current->val = $sum % 10;
            } else {
                $tmp = $current;
                $tmp->next = new ListNode($sum % 10);

                $current = $tmp->next;
            }

            $l1 = $l1->next ?: null;
            $l2 = $l2->next ?: null;
        } while($l1 || $l2 || $extra);
        
        return $root;
        // $dummy_head = new ListNode();
        // $res = $dummy_head;
        // $carry = 0;
        // while ($l1 != null || $l2 != null){
        //     $sum = $carry;
        //     $carry = 0;
        //     if ($l1 !== null) {
        //         $sum += $l1->val;
        //         $l1 = $l1->next;
        //     }
        //     if ($l1 !== null){
        //         $sum += $l2->val;
        //         $l2 = $l2->next;
        //     }

        //     if ($sum > 9){
        //         $carry++;
        //         $sum = $sum % 10;
        //     }
        //     $res->next = new ListNode($sum);
        //     $res = $res->next;
        // }
        // if ($carry > 0) {
        //     $res->next = new ListNode($carry);
        // }
        
        
        // return $dummy_head->next;
    }
}