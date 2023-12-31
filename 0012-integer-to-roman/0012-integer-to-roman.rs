impl Solution {
    pub fn int_to_roman(num: i32) -> String {        
        let (
            one,
            ten,
            hund,
            thou
        ) = (
            ["", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"],
            ["", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"],
            ["", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"],
            ["", "M", "MM", "MMM"]
        );
        
        let num = num as usize;
        format!("{}{}{}{}", 
            thou[num / 1000],
            hund[(num % 1000) / 100],
            ten[(num % 100) / 10],
            one[num % 10]
        )
    }
}