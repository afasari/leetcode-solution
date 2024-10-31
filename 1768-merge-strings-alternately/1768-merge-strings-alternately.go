func mergeAlternately(word1 string, word2 string) string {
//     l1 := len(word1)
//     l2 := len(word2)
    
//     var ret string
    
//     isCheckInside := l1 > l2
//     for i, w := range word1{
//         if isCheckInside && i == l2 {
//             ret += word1[i:l1]
//             break
//         }
        
//         ret += string(w) + string(word2[i])
//     }
    
//     if !isCheckInside{
//         ret += word2[l1:l2]
//     }
    
//     return ret
    
    
    // var res []byte
    var res string
    for i:= 0; i < len(word1) || i < len(word2); i++ {
        if i < len(word1) {
            res += string(word1[i])
        }
        // else {
        //     res = append(res,)
        // }
        
        if i < len(word2) {
            res += string(word2[i])
        } else {
            
        }
    }
    
    return res
    // return handleMerge(word1, word2)
}


// func handleMerge(w1, w2 string) string {
//     var ret string
    
//     isCheckInside := len(w1) > len(w2)
//     for i, w := range w1{
//         if isCheckInside && i == len(w2) {
//             ret += w1[i:len(w1)]
//             break
//         }
        
//         ret += string(w) + string(w2[i])
//     }
    
//     if !isCheckInside{
//         ret += w2[len(w1):len(w2)]
//     }
    
    
//     return ret
// }