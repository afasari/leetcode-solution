class Solution {

    /**
     * @param String $s
     * @param Integer $numRows
     * @return String
     */
    function convert($s, $numRows) {
        $length = strlen($s);
        if($length <= 1 || $numRows <= 1){
            return $s;
        }
        $two_dimensional_array = [];
        $index_x = 0;
        $index_y = 0;
        $array_s = str_split($s);
        $status = "down";
        for($k = 0; $k < $length; $k ++){
            $two_dimensional_array[$index_x][$index_y] = $array_s[$k];
            if("down" == $status){
                $index_x ++;
            }
            if("up_right" == $status){
                $index_x --;
                $index_y ++;
            }
            if($numRows == $index_x){
                $index_x -= 2;
                $index_y ++;
                $status = "up_right";
            }
            if(-1 == $index_x){
                $index_x += 2;
                $index_y --;
                $status = "down";
            }
        }
        foreach ($two_dimensional_array as $k => $v) {
            $two_dimensional_array[$k] = implode("", $v);
        }
        return implode("", $two_dimensional_array);
    }
}