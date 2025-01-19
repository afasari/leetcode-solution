import "strings"

func reverseWords(s string) string {
    strs := strings.Split(s, " ")

    ret := []string{}
    for i := len(strs)-1; i >= 0; i--{
        if strings.TrimSpace(strs[i]) != ""{
        ret = append(ret, strs[i])
        } 
    }

    return strings.Join(ret, " ")
}