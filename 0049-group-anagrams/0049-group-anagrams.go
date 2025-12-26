func groupAnagrams(strs []string) [][]string {
    // 1. Map to group words by their character count signature
    // Key: [26]int (count of each char a-z), Value: slice of original words
    groups := make(map[[26]int][]string)

    for _, s := range strs {
        // 2. Generate the "signature" for the current word
        var count [26]int
        for i := 0; i < len(s); i++ {
            count[s[i]-'a']++
        }
        
        // 3. Group words sharing the same signature
        groups[count] = append(groups[count], s)
    }

    // 4. Collect all groups into a list of lists
    result := make([][]string, 0, len(groups))
    for _, group := range groups {
        result = append(result, group)
    }
    
    return result
}