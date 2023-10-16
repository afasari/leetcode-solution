class Solution {

    /**
     * @param Integer $num
     * @return String
     */
    function intToRoman($num) {
                $map = [
            1000 => 'M',
            900 => 'CM',
            500 => 'D',
            400 => 'CD',
            100 => 'C',
            90 => 'XC',
            50 => 'L',
            40 => 'XL',
            10 => 'X',
            9 => 'IX',
            5 => 'V',
            4 => 'IV',
            1 => 'I',
        ];
        
        $str = '';
        while ($num > 0) {
            foreach($map as $mn => $s) {
                if ($mn <= $num) {
                    $str .= $s;
                    $num -= $mn;
                    break;
                }
            }
        }
        
        return $str;
    }
}