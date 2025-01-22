func isValidSudoku(board [][]byte) bool {
    mRow := make(map[int]map[string]bool,9)
    mCol := make(map[int]map[string]bool,9)
    mSqr := make(map[string]map[string]bool,9)

    for a:=0; a < 9; a++{
        mRow[a] = make(map[string]bool, 9)
        mCol[a] = make(map[string]bool, 9)        
        // mSqr[fmt.Sprintf("%d,%d", a, a)] = make(map[string]bool, 9)
    }
    
        mSqr["0,0"] = make(map[string]bool, 9)
        mSqr["0,1"] = make(map[string]bool, 9)
        mSqr["0,2"] = make(map[string]bool, 9)
        mSqr["1,0"] = make(map[string]bool, 9)
        mSqr["1,1"] = make(map[string]bool, 9)
        mSqr["1,2"] = make(map[string]bool, 9)
        mSqr["2,0"] = make(map[string]bool, 9)
        mSqr["2,1"] = make(map[string]bool, 9)
        mSqr["2,2"] = make(map[string]bool, 9)

    for i := 0; i < 9; i++{
        for j := 0; j < 9; j++{
            val := string(board[i][j])
            if val == "."{
                continue
            }

            idxSqr := fmt.Sprintf("%d,%d",i/3,j/3)
            if mRow[i][val] || mCol[j][val] || mSqr[idxSqr][val]{
                return false
            }
            mRow[i][val] = true
            mCol[j][val] = true
            mSqr[idxSqr][val] = true

        }
    }
    
    return true
}