func selfDividingNumbers(left int, right int) []int {
    ret := []int{}

    for left <= right {
        if isSelfDriving(left) {
            ret = append(ret, left)
        }
        left++
    }


    return ret
}

func isSelfDriving(num int) bool {
    str := strconv.Itoa(num)
    for _, s := range str {
        digit := int(s - '0')

        if digit == 0 || num % digit != 0 {
            return false
        }
    }
    return true
}