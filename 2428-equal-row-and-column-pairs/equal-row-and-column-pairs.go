// Approach 1: Brute-force
func equalPairs(grid [][]int) int {
	n := len(grid)
	count := 0

	reversedGrid := [][]int{}
	for i := range n {
		col := []int{}
		for j := range n {
			col = append(col, grid[j][i])
		}
		reversedGrid = append(reversedGrid, col)
	}

	for i := range n {
		for j := range n {
			if isEqual(grid[i], reversedGrid[j]) {
				count++
			}
		}
	}

	return count
}

func isEqual(row []int, col []int) bool {
	n := len(row)
	for i := range n {
		if row[i] != col[i] {
			return false
		}
	}

	return true
}
