func generate(numRows int) [][]int {
    var result [][]int
	if numRows >= 1 {
		result = append(result, []int{1})
	}

	for i := 1; i < numRows; i++ {
		row := make([]int, i+1)
		row[0] = 1
		row[i] = 1
		for j := 1; j < i; j++ {
			row[j] = result[i-1][j-1] + result[i-1][j]
		}
		result = append(result, row)
	}

	return result
}