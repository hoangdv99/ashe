func convertToTitle(columnNumber int) string {
    result := ""
    for columnNumber > 0 {
        remain := (columnNumber - 1) % 26
        result = string('A' + remain) + result
        columnNumber = (columnNumber - 1) / 26
    }
    return result
}