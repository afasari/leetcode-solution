/** 
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is higher than the picked number
 *			      1 if num is lower than the picked number
 *               otherwise return 0
 * func guess(num int) int;
 */

func guessNumber(n int) int {
    // return sort.Search(n, func (i int) bool {
    //     return guess(1+i) == -1
    // })

	start, end := 1, n
	for start <= end {
		mid := (start + end) / 2
		switch guess(mid) {
		case -1:
			end = mid - 1
		case 1:
			start = mid + 1
		default:
			return mid
		}
	}
    return 0
    
    l, r, m := 1, n, 0
    for l <= r{
        m = (l + r) >> 1
        if g:= guess(m); g == -1{
            r = m-1
        } else if g == 1{
            l = m + 1
        } else {
            break
        }
    }

    return m
    
}