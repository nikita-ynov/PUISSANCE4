package utils

import "power4/controller/structure"

func PlacePiece(piece int, color string, table *structure.Table) {
	var x, y int = 0, 0
	switch piece {
	case 1:
		x = 0
	case 2:
		x = 70
	case 3:
		x = 70 * 2
	case 4:
		x = 70 * 3
	case 5:
		x = 70 * 4
	case 6:
		x = 70 * 5
	}
	table.Placement = append(table.Placement, structure.Placement{X: x, Y: y, Color: color})
}
