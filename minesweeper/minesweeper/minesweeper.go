package minesweeper

import "math/rand"

type Cell struct {
	I, J int
}

type Minesweeper struct {
	height, width, totalMines int
	Board                     [][]bool
	GameOver, Lost            bool
	minesSet                  map[Cell]bool
	fieldsSet                 map[Cell]bool
}

func NewGame(height, width, mines int) *Minesweeper {
	game := Minesweeper{
		height:     height,
		width:      width,
		totalMines: mines,
		Board:      make([][]bool, height),
		minesSet:   make(map[Cell]bool),
		fieldsSet:  make(map[Cell]bool),
	}

	for i := 0; i < height; i++ {
		game.Board[i] = make([]bool, width)
	}

	game.addMines()

	return &game
}

func (m *Minesweeper) addMines() {
	count := 0
	for {
		if count == m.totalMines {
			break
		}

		i := rand.Intn(m.height)
		j := rand.Intn(m.width)

		if m.Board[i][j] {
			continue
		}

		m.Board[i][j] = true
		m.minesSet[Cell{i, j}] = true
		count++
	}
}

func (m *Minesweeper) IsValidCell(i, j int) bool {
	return 0 <= i && i < m.height && 0 <= j && j < m.width
}

func (m *Minesweeper) IsMine(cell Cell) bool {
	return m.Board[cell.I][cell.J]
}

func (m *Minesweeper) NearbyMines(cell Cell) int {
	count := 0

	for i := cell.I - 1; i <= cell.I+1; i++ {
		for j := cell.J - 1; j <= cell.J+1; j++ {
			if !m.IsValidCell(i, j) || (i == cell.I && j == cell.J) {
				continue
			}

			if m.Board[i][j] {
				count++
			}
		}
	}

	return count
}

func (m *Minesweeper) IsGameOver() bool {
	if m.Lost || m.GameOver {
		return true
	}

	if len(m.fieldsSet) == m.height*m.width {
		return true
	}

	return false
}

func (m *Minesweeper) Won(minesFound map[Cell]bool) bool {
	if len(m.minesSet) != len(minesFound) {
		return false
	}

	for cell := range m.minesSet {
		if !minesFound[cell] {
			return false
		}
	}
	return true
}
