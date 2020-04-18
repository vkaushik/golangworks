package problem/rottingOranges

func orangesRotting(grid [][]int) int {
	freshOranges := 0
	rottenOranges := 0
	minutesToRotAll := 0

	for row := range grid {
		for column := range grid[row] {
			if orange := grid[row][column]; orange == 1 {
				freshOranges++
			} else if orange == 2 {
				rottenOranges++
			}
		}
	}

	for rottenOranges > 0 && freshOranges > 0 {
		copyGrid := createCopyGrid(grid)
		isSomethingRotting := false

		for row := range grid {
			for column := range grid[row] {
				if isRotting := isStartedRotting(row, column, copyGrid); isRotting {
					grid[row][column] = 2
					rottenOranges++
					freshOranges--
					isSomethingRotting = true
				}
			}
		}

		if !isSomethingRotting {
			break
		}

		minutesToRotAll++
	}

	if freshOranges != 0 {
		return -1
	}

	return minutesToRotAll
}

func createCopyGrid(grid [][]int) [][]int {
	copyGrid := make([][]int, len(grid))

	for rowIndex, row := range grid {
		copyGrid[rowIndex] = make([]int, len(row))
		copy(copyGrid[rowIndex], row)
	}

	return copyGrid
}

func isStartedRotting(row int, column int, grid [][]int) bool {
	orange := grid[row][column]
	if orange == 0 || orange == 2 {
		return false
	}

	return hasAnyAdjacentRottenOrange(row, column, grid)
}

func hasAnyAdjacentRottenOrange(row int, column int, grid [][]int) bool {

	return hasRottingOrangeOnLeft(row, column, grid) ||
		hasRottingOrangeOnRight(row, column, grid) ||
		hasRottingOrangeOnTop(row, column, grid) ||
		hasRottingOrangeOnBottom(row, column, grid)
}

func hasRottingOrangeOnLeft(row int, column int, grid [][]int) bool {
	// column--
	return column > 0 && grid[row][column-1] == 2
}

func hasRottingOrangeOnRight(row int, column int, grid [][]int) bool {
	// column++
	lastColumnIndex := len(grid[0]) - 1
	return column < lastColumnIndex && grid[row][column+1] == 2
}

func hasRottingOrangeOnTop(row int, column int, grid [][]int) bool {
	// row--
	return row > 0 && grid[row-1][column] == 2
}

func hasRottingOrangeOnBottom(row int, column int, grid [][]int) bool {
	// row++
	lastRowIndex := len(grid) - 1
	return row < lastRowIndex && grid[row+1][column] == 2
}
