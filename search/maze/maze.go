package maze

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"slices"

	"github.com/fogleman/gg"

	"github.com/samuelralmeida/ai-topics/search/challenges"
	"github.com/samuelralmeida/ai-topics/search/entity"
)

type Frontier interface {
	Add(node entity.Node)
	IsEmpty() bool
	Remove() entity.Node
	ContainsState(state entity.Coordinate) bool
}

type Maze struct {
	Height int
	Width  int
	Walls  [][]string
	Start  entity.Coordinate
	Goal   string
}

func NewMaze(filename string, goalChar string) (*Maze, error) {
	maze := Maze{}

	err := maze.buildWalls(filename)
	if err != nil {
		return nil, err
	}

	start := entity.Coordinate{}

	for i, row := range maze.Walls {
		for j, collumn := range row {
			if collumn == "A" {
				start.Collumn = j
				start.Row = i
			}
		}
	}

	maze.Start = start
	maze.Goal = goalChar

	maze.Height = len(maze.Walls) - 1
	maze.Width = len(maze.Walls[0]) - 1

	return &maze, nil
}

func (m *Maze) possibleActions(state entity.Coordinate) []entity.Action {
	type moveCandidate struct {
		action     string
		coordinate entity.Coordinate
	}

	candidates := []moveCandidate{
		{action: "up", coordinate: entity.Coordinate{Row: state.Row - 1, Collumn: state.Collumn}},
		{action: "down", coordinate: entity.Coordinate{Row: state.Row + 1, Collumn: state.Collumn}},
		{action: "left", coordinate: entity.Coordinate{Row: state.Row, Collumn: state.Collumn - 1}},
		{action: "right", coordinate: entity.Coordinate{Row: state.Row, Collumn: state.Collumn + 1}},
	}

	actions := []entity.Action{}
	for _, candidate := range candidates {
		if m.isValidCoordinate(candidate.coordinate) && !m.isWall(candidate.coordinate) {
			actions = append(actions, entity.Action{Description: candidate.action, Row: candidate.coordinate.Row, Collumn: candidate.coordinate.Collumn})
		}
	}
	return actions
}

func (m *Maze) isValidCoordinate(coordinate entity.Coordinate) bool {
	return coordinate.Row >= 0 && coordinate.Row <= m.Height && coordinate.Collumn >= 0 && coordinate.Collumn <= m.Width
}

func (m *Maze) isWall(coordinate entity.Coordinate) bool {
	return m.Walls[coordinate.Row][coordinate.Collumn] == "#"
}

func (m *Maze) buildWalls(filename string) error {
	file, err := challenges.FS.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	walls := [][]string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		row := []string{}
		for _, char := range line {
			element := string(char)
			if element == "" {
				element = " "
			}
			row = append(row, element)
		}
		walls = append(walls, row)
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	m.Walls = walls

	return nil
}

func (m *Maze) Solve(frontier Frontier) (*entity.Solution, error) {
	numExplored := 0
	explored := make(map[entity.Coordinate]bool)

	start := entity.Node{State: m.Start}
	frontier.Add(start)

	for {

		if frontier.IsEmpty() {
			return nil, errors.New("no solution")
		}

		node := frontier.Remove()
		numExplored++
		if m.Walls[node.State.Row][node.State.Collumn] == m.Goal {
			actions := []entity.Action{}
			cells := []entity.Coordinate{}

			for node.Parent != nil {
				actions = append(actions, node.Action)
				cells = append(cells, node.State)
				node = *node.Parent
			}

			slices.Reverse(actions)
			slices.Reverse(cells)

			return &entity.Solution{Actions: actions, Cells: cells, NumExplored: numExplored, NodesExplored: explored}, nil
		}

		explored[node.State] = true
		for _, action := range m.possibleActions(node.State) {
			state := entity.Coordinate{Row: action.Row, Collumn: action.Collumn}
			_, wasExĺored := explored[state]
			if !frontier.ContainsState(state) && !wasExĺored {
				child := entity.Node{State: state, Parent: &node, Action: action}
				frontier.Add(child)
			}
		}

	}
}

func (m *Maze) PrintSolve(frontier Frontier) {
	solution, err := m.Solve(frontier)
	if err != nil {
		fmt.Println(err)
		return
	}

	maze := m.Walls
	cells := solution.Cells[:len(solution.Cells)-1]

	for _, coordinate := range cells {
		maze[coordinate.Row][coordinate.Collumn] = "*"
	}

	for _, row := range maze {
		for _, collumn := range row {
			fmt.Printf("%s ", collumn)
		}
		fmt.Println("")
	}
	fmt.Println("path cost:", solution.NumExplored)
}

func (m *Maze) ImageSolve(filename string, frontier Frontier, showSolution, showExplored bool) {
	solution, err := m.Solve(frontier)
	if err != nil {
		fmt.Println(err)
		return
	}

	cellSize := 50.0
	cellBorder := 2.0

	// Create a new context
	dc := gg.NewContext(int(cellSize*float64(m.Width+1)), int(cellSize*float64(m.Height+1)))
	dc.SetRGB255(0, 0, 0)
	dc.Clear()

	for i, row := range m.Walls {
		for j, col := range row {
			coordinate := entity.Coordinate{Row: i, Collumn: j}

			if col == "#" {
				// walls
				dc.SetRGB255(40, 40, 40)
			} else if coordinate == m.Start {
				// start
				dc.SetRGB255(255, 0, 0)
			} else if col == m.Goal {
				// goal
				dc.SetRGB255(0, 171, 28)
			} else if showSolution && slices.Contains(solution.Cells, coordinate) {
				// solution
				dc.SetRGB255(220, 235, 113)
			} else if showExplored && solution.NodesExplored[coordinate] {
				// explored
				dc.SetRGB255(212, 97, 85)
			} else {
				// empty cell
				dc.SetRGB255(237, 240, 252)
			}

			// draw cell
			dc.DrawRectangle(
				float64(j)*cellSize+cellBorder, float64(i)*cellSize+cellBorder,
				cellSize-2*cellBorder, cellSize-2*cellBorder,
			)
			dc.Fill()
		}
	}

	// Save the image to a file
	err = dc.SavePNG(filename)
	if err != nil {
		log.Fatal(err)
	}
}
