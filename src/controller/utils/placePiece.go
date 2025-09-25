package utils

import "power4/controller/structure"

func PlacePiece(piece string, color string, table *structure.Table) *structure.Table {
	var x int

	switch piece {
	case "1":
		x = 0
	case "2":
		x = 70
	case "3":
		x = 70 * 2
	case "4":
		x = 70 * 3
	case "5":
		x = 70 * 4
	case "6":
		x = 70 * 5
	default:
		return table // pièce invalide → on ne change rien
	}

	y := 0 // TODO : calculer en fonction de la gravité (première ligne libre)

	table.Placement = append(table.Placement, structure.Placement{
		X:     x,
		Y:     y,
		Color: color,
	})

	return table
}
