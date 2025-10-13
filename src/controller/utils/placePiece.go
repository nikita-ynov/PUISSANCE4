package utils

import "power4/controller/structure"

func PlacePiece(piece string, color string, table *structure.Table) (*structure.Table, bool) {
	const (
		cellHeight = 70
		rows       = 6
		columns    = 7
	)
	var xIndex int
	switch piece {
	case "1":
		xIndex = 0
	case "2":
		xIndex = 1
	case "3":
		xIndex = 2
	case "4":
		xIndex = 3
	case "5":
		xIndex = 4
	case "6":
		xIndex = 5
	case "7":
		xIndex = 6
	default:
		return table, false
	}

	// Compter combien de pions sont déjà dans cette colonne
	targetX := xIndex * cellHeight
	count := 0
	maxY := 0
	for _, p := range table.Placement {
		if p.X == targetX {
			count++
			if p.Y >= maxY {
				maxY = p.Y
			}
		}
	}

	// Colonne pleine (6 rangées)
	if count >= rows {
		return table, false
	}

	// Positionner le pion à la première case libre (gravité)
	y := maxY
	if count > 0 {
		y = maxY + cellHeight
	} else {
		y = 0
	}

	table.Placement = append(table.Placement, structure.Placement{
		X:     targetX,
		Y:     y,
		Color: color,
	})

	return table, true
}
