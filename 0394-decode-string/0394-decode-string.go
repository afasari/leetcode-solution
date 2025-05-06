func decodeString(s string) string {
	rs := []byte{}

	for i := range s {
		if s[i] != ']' {
			rs = append(rs, s[i])
			continue
		}

		j := len(rs) - 1
		encStr := []byte{}
		for ; rs[j] != '['; j-- {
			encStr = append(encStr, rs[j])
		}

		k := j - 1
		rep := []byte{}
		for ; k >= 0 && rs[k] >= '0' && rs[k] <= '9'; k-- {
			rep = append(rep, rs[k])
		}
		slices.Reverse(rep)
		repNum, _ := strconv.Atoi(string(rep))

		if k >= 0 {
			rs = append(rs[:k + 1], rs[j+1:]...)
		} else {
			rs = rs[j+1:]
		}

		for occur := 1; occur < repNum; occur++ {
			for encIdx := len(encStr) - 1; encIdx >= 0; encIdx-- {
                rs = append(rs, encStr[encIdx])
			}
		}
	}
	return string(rs)
}