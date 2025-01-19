import "strings"

func reverseWords(s string) string {
    strs := strings.Split(s, " ")

    ret := []string{}
    for i := len(strs)-1; i >= 0; i--{
        if strings.TrimSpace(strs[i]) != ""{
        ret = append(ret, strs[i])
fmt.Println(strs[i])
        } 
    }

// fmt.Println(strs)
    return strings.Join(ret, " ")

    
}