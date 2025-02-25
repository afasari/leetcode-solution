/** 
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is higher than the picked number
 *			      1 if num is lower than the picked number
 *               otherwise return 0
 * func guess(num int) int;
 */

func guessNumber(n int) int {
    return sort.Search(n, func (i int) bool {
        return guess(1+i) == -1
    })
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