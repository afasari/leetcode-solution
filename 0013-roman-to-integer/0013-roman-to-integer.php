class Solution {

    /**
     * @param String $s
     * @return Integer
     */
    function romanToInt($s) {
                $dict = [
            'I'  => 1,
            'IV' => 4,
            'V'  => 5,
            'IX' => 9,
            'X'  => 10,
            'XL' => 40,
            'L'  => 50,
            'XC' => 90,
            'C'  => 100,
            'CD' => 400,
            'D'  => 500,
            'CM' => 900,
            'M'  => 1000,
        ];
        $s=strtoupper($s);
        if(isset($dict[$s])){
            return $dict[$s];
        }
        $len=strlen($s);
        $value=0;
        for($i=0;$i<$len;$i++){
            $less=$s[$i];
            $value+=$dict[$less];
            if(($i-1)>=0) {
                $fLess=$s[$i-1];
                if($fLess && $dict[$fLess]<$dict[$less]){
                    $value-=($dict[$fLess]*2);
                }
            }
        }

        return $value;
    }
}