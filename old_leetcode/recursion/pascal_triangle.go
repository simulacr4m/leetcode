
func getRowHelper(rowIndex int, target int, rows [][]int) [][]int {
    if rowIndex > target {
        return rows
    } else if rowIndex == 0 {
        rows = append(rows, []int{1})
        rows = getRowHelper(rowIndex+1, target, rows)
        return rows
    } else if rowIndex == 1 {
        rows = append(rows, []int{1, 1})
        rows = getRowHelper(rowIndex+1, target, rows)
        return rows
    }
    
    var row []int = []int{1}
    for i := 0; i < len(rows[rowIndex-1])-1; i++ {
        row = append(row, rows[rowIndex-1][i] + rows[rowIndex-1][i+1])
    }
    row = append(row, 1)
    
    rows = append(rows, row)
    rows = getRowHelper(rowIndex+1, target, rows)
    return rows
}

func getRow(rowIndex int) []int {
    var rows [][]int
    rows = getRowHelper(0, rowIndex, rows)
    return rows[len(rows)-1]
}
