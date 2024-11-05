func gcdOfStrings(str1 string, str2 string) string {
    if str1 + str2 != str2 + str1 {
        return ""
    }
    
    x := gcd(len(str1), len(str2))
    return str1[:x]
}


func gcd(a, b int) int {
    for b != 0 {
        a, b = b, a % b
    }
    return a
}

// func gcdOfStrings(str1 string, str2 string) string {

    
//     if str1 < str2{
//         str1, str2 = str2, str1
//     }
//     l1 := len(str1)
//     l2 := len(str2)
    
//     if str1 == str2{
//         fmt.Println("masuk sama dengan")
//         return str1
//     }
    
//     if l1 % l2 == 0{
//         fmt.Println("masuk modulus")
//         for i :=0; i < l1; i += l2{
//             if string(str1[i:i+l2])!= str2{
//                 return ""
//             }
//         }
        
//         return str2
//     }
    
//     fmt.Println("masuk luar")

//     res := ""
//     for i, r := range str2{
//         if string(r) == string(str1[i]){
//             res += string(r)
//             lres := len(res)
//             if l1 % lres != 0 || l2 % lres != 0{
//                 return res[0:lres-1]
//             }
//         } else {
//             return ""
//         }
//     }

//     return res
// }