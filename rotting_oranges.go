package main

func rottingOranges(grid [][]int) int {
	isFreshOranges := true
	mins := 0
	for isFreshOranges {
		change := false
		mins++
		for i, row := range grid {
			for j, column := range row {
				if column == 1 {
					if (0 <= i-1 && i-1 < len(grid) && grid[i-1][j] >= 2 && grid[i-1][j] < 2+mins) ||
						(0 <= i+1 && i+1 < len(grid) && grid[i+1][j] >= 2 && grid[i+1][j] < 2+mins) ||
						(0 <= j-1 && j-1 < len(row) && grid[i][j-1] >= 2 && grid[i][j-1] < 2+mins) ||
						(0 <= j+1 && j+1 < len(row) && grid[i][j+1] >= 2 && grid[i][j+1] < 2+mins) {
						grid[i][j] = 2 + mins
						change = true
					}
				}
			}
		}
		if !change {
			isFreshOranges = false
		}
	}

	return mins - 1
}
