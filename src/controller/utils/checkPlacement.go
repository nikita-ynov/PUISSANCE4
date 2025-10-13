package utils

import "power4/controller/structure"

const (
	columns  = 7
	rows     = 6
	cellSize = 70
)

func CheckPlacement(table *structure.Table) string {
	grid := [rows][columns]string{}

	// remplir la grille
	for _, p := range table.Placement {
		col := p.X / cellSize
		row := p.Y / cellSize
		if row < rows && col < columns {
			grid[row][col] = p.Color
		}
	}

	// vérifier les directions
	for r := 0; r < rows; r++ {
		for c := 0; c < columns; c++ {
			color := grid[r][c]
			if color == "" {
				continue
			}

			// horizontal
			if c+3 < columns &&
				grid[r][c+1] == color &&
				grid[r][c+2] == color &&
				grid[r][c+3] == color {
				return color
			}
			// vertical
			if r+3 < rows &&
				grid[r+1][c] == color &&
				grid[r+2][c] == color &&
				grid[r+3][c] == color {
				return color
			}
			// diagonale /
			if r+3 < rows && c-3 >= 0 &&
				grid[r+1][c-1] == color &&
				grid[r+2][c-2] == color &&
				grid[r+3][c-3] == color {
				return color
			}
			// diagonale \
			if r+3 < rows && c+3 < columns &&
				grid[r+1][c+1] == color &&
				grid[r+2][c+2] == color &&
				grid[r+3][c+3] == color {
				return color
			}
		}
	}
	return ""
}
