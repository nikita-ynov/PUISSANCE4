package structure

type Table struct {
	Placement []Placement
}

type Placement struct {
	x     int
	y     int
	color string
}
