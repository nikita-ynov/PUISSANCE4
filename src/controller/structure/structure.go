package structure

type Table struct {
	Placement []Placement
}

type Placement struct {
	X     int
	Y     int
	Color string
}
