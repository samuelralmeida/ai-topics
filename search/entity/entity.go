package entity

type Coordinate struct {
	Row     int
	Collumn int
}

type Action struct {
	Description string
	Row         int
	Collumn     int
}

type Solution struct {
	Actions       []Action
	Cells         []Coordinate
	NumExplored   int
	NodesExplored map[Coordinate]bool
}

type Node struct {
	State  Coordinate
	Parent *Node
	Action Action
}
