package frontier

import (
	"slices"

	"github.com/samuelralmeida/ai-topics/search/entity"
)

type aStarFrontier struct {
	frontier []entity.Node
	goal     entity.Coordinate
}

// NewAStarFrontier returns aStar frontier which applies A* search algorithm
func NewAStarFrontier(goal entity.Coordinate) *aStarFrontier {
	return &aStarFrontier{frontier: []entity.Node{}, goal: goal}
}

func (af *aStarFrontier) Add(node entity.Node) {
	af.frontier = append(af.frontier, node)
}

func (af *aStarFrontier) IsEmpty() bool {
	return len(af.frontier) == 0
}

func (af *aStarFrontier) Remove() entity.Node {
	if af.IsEmpty() {
		return entity.Node{}
	}

	var lowerCost int
	var indexNode int

	for i, node := range af.frontier {
		cost := af.manhatanHeuristic(node.State) + node.ReachCost
		if lowerCost == 0 || cost < lowerCost {
			lowerCost = cost
			indexNode = i
		}
	}

	node := af.frontier[indexNode]
	af.frontier = slices.Delete(af.frontier, indexNode, indexNode+1)
	return node
}

func (af *aStarFrontier) ContainsState(state entity.Coordinate) bool {
	for _, node := range af.frontier {
		if node.State == state {
			return true
		}
	}
	return false
}

func (af *aStarFrontier) manhatanHeuristic(state entity.Coordinate) int {
	rowDiff := af.goal.Row - state.Row
	if rowDiff < 0 {
		rowDiff = rowDiff * (-1)
	}

	colDiff := af.goal.Collumn - state.Collumn
	if colDiff < 0 {
		colDiff = colDiff * (-1)
	}

	return rowDiff + colDiff
}
