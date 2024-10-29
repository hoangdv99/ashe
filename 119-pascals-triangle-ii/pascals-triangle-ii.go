func getRow(rowIndex int) []int {
    if rowIndex == 0 {
		return []int{1}
	}
	var triangle [][]int
	if rowIndex >= 1 {
		triangle = append(triangle, []int{1})
	}

	for i := 1; i <= rowIndex; i++ {
		row := make([]int, i+1)
		row[0] = 1
		row[i] = 1
		for j := 1; j < i; j++ {
			row[j] = triangle[i-1][j-1] + triangle[i-1][j]
		}
		triangle = append(triangle, row)
	}

	return triangle[rowIndex]
}