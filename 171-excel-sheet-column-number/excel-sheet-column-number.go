func titleToNumber(columnTitle string) int {
    columnNumber := 0
    for i := 0; i < len(columnTitle); i++ {
        columnNumber = columnNumber * 26 + int(columnTitle[i] - 'A' + 1)
    }
    return columnNumber
}