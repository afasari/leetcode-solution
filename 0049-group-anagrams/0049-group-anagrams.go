func groupAnagrams(strs []string) [][]string {

    mapSortedStr := make(map[string][]string)


    for _, str := range strs{
        strSplit := strings.Split(str, "")
        sort.Strings(strSplit)
        sortedStr := strings.Join(strSplit, "")
        mapSortedStr[sortedStr] = append(mapSortedStr[sortedStr], str)
    }

    var res [][]string

    for _, val := range mapSortedStr {
        res = append(res, val)
    }

    return res
}