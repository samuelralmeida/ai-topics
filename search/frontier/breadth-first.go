package frontier

import (
	"github.com/samuelralmeida/ai-topics/search/entity"
)

type queueFrontier struct {
	frontier []entity.Node
}

// NewQueuerFrontier returns queue frontier which applies bread-first search algorithm using queue data structure
func NewQueueFrontier() *queueFrontier {
	return &queueFrontier{frontier: []entity.Node{}}
}

func (qf *queueFrontier) Add(node entity.Node) {
	qf.frontier = append(qf.frontier, node)
}

func (qf *queueFrontier) IsEmpty() bool {
	return len(qf.frontier) == 0
}

func (qf *queueFrontier) Remove() entity.Node {
	if qf.IsEmpty() {
		return entity.Node{}
	}

	node := qf.frontier[0]
	qf.frontier = qf.frontier[1:]
	return node
}

func (qf *queueFrontier) ContainsState(state entity.Coordinate) bool {
	for _, node := range qf.frontier {
		if node.State == state {
			return true
		}
	}
	return false
}
