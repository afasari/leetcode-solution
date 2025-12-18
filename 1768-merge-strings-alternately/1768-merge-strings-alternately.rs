impl Solution {
     pub fn merge_alternately(word1: String, word2: String) -> String {
        let k = word1.len() + word2.len();
        let mut merged = String::with_capacity(k);
        let (mut word1, mut word2) = (word1.chars(), word2.chars());
        
        for _ in 0..k {
            if let Some(c) = word1.next() {
                merged.push(c);
            }
            if let Some(c) = word2.next() {
                merged.push(c);
            }
        }
        merged
    }
}