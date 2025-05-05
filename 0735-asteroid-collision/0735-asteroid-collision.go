func asteroidCollision(astr []int) []int {
    prev := -1
    for i := 0; i < len(astr); i++ {
        notSame := true
        if astr[i] > 0 {
            prev++
            astr[prev] = astr[i]
        } else {
            for prev >= 0 && astr[prev] > 0 {
                if astr[prev] < -astr[i] {
                    prev--
                } else if astr[prev] == -astr[i] {
                    prev--
                    notSame = false
                    break
                } else {
                    break
                }
            }

            if notSame && (prev < 0 || astr[prev] < 0) {
                prev++
                astr[prev] = astr[i]
            }
        }
    }
    return astr[:prev+1]
}