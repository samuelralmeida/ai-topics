package frontier

import (
	"github.com/samuelralmeida/ai-topics/search/entity"
)

type stackFrontier struct {
	frontier []entity.Node
}

// NewStackFrontier returns stack frontier which applies depth-first search algorithm using stack data structure
func NewStackFrontier() *stackFrontier {
	return &stackFrontier{frontier: []entity.Node{}}
}

func (sf *stackFrontier) Add(node entity.Node) {
	sf.frontier = append(sf.frontier, node)
}

func (sf *stackFrontier) IsEmpty() bool {
	return len(sf.frontier) == 0
}

func (sf *stackFrontier) Remove() entity.Node {
	if sf.IsEmpty() {
		return entity.Node{}
	}

	node := sf.frontier[len(sf.frontier)-1]
	sf.frontier = sf.frontier[:len(sf.frontier)-1]
	return node
}

func (sf *stackFrontier) ContainsState(state entity.Coordinate) bool {
	for _, node := range sf.frontier {
		if node.State == state {
			return true
		}
	}
	return false
}
