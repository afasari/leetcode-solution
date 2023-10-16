class Solution {

    /**
     * @param String $s
     * @return Boolean
     */
    function isValid($s) {
        $stack = [];
        $map = [
            "(" => ")",
            "{" => "}",
            "[" => "]",
        ];
        if (strlen($s) === 1){
            return false;
        }
        if (strlen($s) === 2){
            return $s[1] == $map[$s[0]];
        }

        $split = str_split($s);

        foreach($split as $l){
            if (isset($map[$l])){
                // append to stack
                array_push($stack, $l);
            } else{
                // check on the last stack is paired of the map?
                if ($map[$stack[count($stack)-1]] !== $l){
                    return false;
                }
                array_pop($stack);
            }
        }

        return count($stack) == 0;
        
    }
}