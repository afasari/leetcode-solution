
func romanToInt(s string) int {
	values := map[string]int{
		"I":  1,
		"IV": 4,
		"V":  5,
		"IX": 9,
		"X":  10,
		"XL": 40,
		"L":  50,
		"XC": 90,
		"C":  100,
		"CD": 400,
        "D":  500,
		"CM": 900,
		"M":  1000,
	}
    
    res := 0
    i := 0
    for i < len(s) {
        if i < len(s)-1 && values[s[i:i+2]] != 0 {
            res += values[s[i:i+2]]
            i += 2
        } else {
            res += values[s[i:i+1]]
            i++
        }
    }
    
    return res
}