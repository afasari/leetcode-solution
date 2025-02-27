// func tribonacci(n int) int {
//     triset := [3]int{0, 1, 1}
//     if n < 3 {
//         return triset[n]
//     }
//     for i := 3; i <= n; i++ {
//         triset[i % 3] = triset[0] + triset[1] + triset[2]
//     }
//     return triset[n % 3]
// }

func tribonacci(n int) int {
    // open slice to store elemts
    sequence := make([]int,0,n)
    // add first three elemts so we will use them to form element T_3
    sequence = append(sequence,0,1,1)
    // iterate throught sequence 
    for i:= 3;i<=n; i++ {
    // calculate the summ of three privous elements
        sum:=sequence[i-1] + sequence[i-2] + sequence[i-3] 
    // append the sum to sequence so it be our next element  
        sequence = append(sequence,sum)
    }
    // return last element of slice because we need value of nth Tribonacci
    return sequence[n]
}