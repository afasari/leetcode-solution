func equalPairs(grid [][]int) int {
    rowMap := map[string]int{}
    colMap := map[string]int{}
    count := 0

    for i:=0; i<len(grid); i++{
        col := ""
        row := ""
        for j:=0; j<len(grid); j++{
            row += strconv.Itoa(grid[i][j]) + ","
            col += strconv.Itoa(grid[j][i]) + ","
        }
        rowMap[row]++
        colMap[col]++
    }

    for key, val := range(rowMap){
        count += val * colMap[key]
    }

    return count
}