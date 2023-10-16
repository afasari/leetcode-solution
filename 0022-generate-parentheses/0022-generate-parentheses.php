class Solution {

    /**
     * @param Integer $n
     * @return String[]
     */

    protected $result = [];
    protected $length = 0;
    function generateParenthesis($n) {
        $this->length = $n * 2;
        $this->backtracking('', 0, 0, $n);
        return $this->result;
    }

    function backtracking($parenthesis, $left, $right, $n) {
        if (strlen($parenthesis) == $this->length) { 
            array_push($this->result, $parenthesis);
        }
        if ($left < $n) {
            $this->backtracking($parenthesis . '(', $left + 1, $right, $n); 
        }
        if ($right < $left) { 
            $this->backtracking($parenthesis . ')', $left, $right + 1, $n); 
        }
    }
}