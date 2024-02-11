package entity

type Cordinate struct {
	Row     int
	Collumn int
}

type Action struct {
	Description string
	Row         int
	Collumn     int
}

type Solution struct {
	Actions     []Action
	Cells       []Cordinate
	NumExplored int
}

type Node struct {
	State  Cordinate
	Parent *Node
	Action Action
}
