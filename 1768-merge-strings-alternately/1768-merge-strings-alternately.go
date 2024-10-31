func mergeAlternately(word1 string, word2 string) string {
    l1 := len(word1)
    l2 := len(word2)
    
    var ret string
    
    if l1 < l2 {
        for i, w := range word1{
            ret += string(w) + string(word2[i])
        }
        
        ret += word2[l1:l2]
    } else {
        for i, w := range word1{
            if i == l2 {
                ret += word1[i:l1]
                break
            }

            ret += string(w) + string(word2[i])
        }
    }

    return ret
}


func handleMerge(w1, w2 string, maxSame int) string {
    var ret string
    for i, w := range w1{
        if i == maxSame {
            fmt.Println(ret)
            fmt.Println(w1[i:len(w1)+1])
            ret += w1[i:len(w1)+1]
            break
        } else {
            fmt.Println("ret", 1, string(w), ret)
        }
        
        ret += string(w) + string(w2[i])
    }
    return ret
}