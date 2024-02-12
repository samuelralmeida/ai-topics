package frontier

import (
	"slices"

	"github.com/samuelralmeida/ai-topics/search/entity"
)

type greedyFrontier struct {
	frontier []entity.Node
	goal     entity.Coordinate
}

// NewGreedyFrontier returns greedy frontier which applies greedy best-first search algorithm
func NewGreedyFrontier(goal entity.Coordinate) *greedyFrontier {
	return &greedyFrontier{frontier: []entity.Node{}, goal: goal}
}

func (gf *greedyFrontier) Add(node entity.Node) {
	gf.frontier = append(gf.frontier, node)
}

func (gf *greedyFrontier) IsEmpty() bool {
	return len(gf.frontier) == 0
}

func (gf *greedyFrontier) Remove() entity.Node {
	if gf.IsEmpty() {
		return entity.Node{}
	}

	var lowerCost int
	var indexNode int

	for i, node := range gf.frontier {
		cost := gf.manhatanHeuristic(node.State)
		if lowerCost == 0 || cost < lowerCost {
			lowerCost = cost
			indexNode = i
		}
	}

	node := gf.frontier[indexNode]
	gf.frontier = slices.Delete(gf.frontier, indexNode, indexNode+1)
	return node
}

func (gf *greedyFrontier) ContainsState(state entity.Coordinate) bool {
	for _, node := range gf.frontier {
		if node.State == state {
			return true
		}
	}
	return false
}

func (gf *greedyFrontier) manhatanHeuristic(state entity.Coordinate) int {
	rowDiff := gf.goal.Row - state.Row
	if rowDiff < 0 {
		rowDiff = rowDiff * (-1)
	}

	colDiff := gf.goal.Collumn - state.Collumn
	if colDiff < 0 {
		colDiff = colDiff * (-1)
	}

	return rowDiff + colDiff
}
