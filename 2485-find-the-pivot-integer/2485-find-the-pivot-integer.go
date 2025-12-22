func pivotInteger(n int) int {
	t := float64(n)
	x := math.Sqrt(t * (t + 1) / 2)
	if z, k := math.Modf(x); k == 0 {
		return int(z)
	}
	return -1    
}