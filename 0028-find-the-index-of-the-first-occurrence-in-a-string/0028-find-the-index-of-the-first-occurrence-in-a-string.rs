impl Solution {
    pub fn str_str(haystack: String, needle: String) -> i32 {
        let lh = haystack.len();
        let ln = needle.len();
        
        if lh < ln{
            return -1
        }
        if lh == ln{
            if haystack == needle{
                return 0;
            }
            return -1;
        }
        for i in 0..=lh-ln {
            let h = &haystack[i..i+ln];
            println!("{i}, {lh}, {ln}, {h}");
            if h == needle{
                return i as i32;
            }
        }
        return -1;
    }
}