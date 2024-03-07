package minesweeper

import (
	"slices"
)

func (c *Cell) Equal(other Cell) bool {
	return c.I == other.I && c.J == other.J
}

type Sentence struct {
	Cells []Cell
	Count int
}

func NewSentence(cells []Cell, count int) Sentence {
	return Sentence{Cells: cells, Count: count}
}

func (s *Sentence) KnownMines() []Cell {
	if len(s.Cells) == s.Count {
		return s.Cells
	}
	return []Cell{}
}

func (s *Sentence) KnownSafes() []Cell {
	if s.Count == 0 {
		return s.Cells
	}
	return []Cell{}
}

func (s *Sentence) MarkMine(cell Cell) {
	index := slices.Index(s.Cells, cell)
	if index >= 0 {
		s.Cells = slices.Delete(s.Cells, index, index+1)
		s.Count--
	}
}

func (s *Sentence) MarkSafe(cell Cell) {
	index := slices.Index(s.Cells, cell)
	if index >= 0 {
		s.Cells = slices.Delete(s.Cells, index, index+1)
	}
}

func (s *Sentence) Equal(other Sentence) bool {
	if len(s.Cells) != len(other.Cells) {
		return false
	}

	if s.Count != other.Count {
		return false
	}

	for _, cell := range s.Cells {
		if !slices.Contains(other.Cells, cell) {
			return false
		}
	}
	return true
}

func (s *Sentence) DifferenceCells(other Sentence) []Cell {
	cells := make([]Cell, 0)
	for _, cell := range s.Cells {
		if !slices.Contains(other.Cells, cell) {
			cells = append(cells, cell)
		}
	}
	return cells
}

func (s *Sentence) IsSubsetCells(other Sentence) bool {
	for _, cell := range s.Cells {
		if !slices.Contains(other.Cells, cell) {
			return false
		}
	}
	return true
}

type MinesweeperAI struct {
	height    int
	width     int
	movesMade map[Cell]bool
	mines     map[Cell]bool
	safes     map[Cell]bool
	knowledge []*Sentence
}

func NewAI(height, width int) *MinesweeperAI {
	return &MinesweeperAI{
		height:    height,
		width:     width,
		movesMade: make(map[Cell]bool),
		mines:     make(map[Cell]bool),
		safes:     make(map[Cell]bool),
		knowledge: make([]*Sentence, 0),
	}
}

func (ai *MinesweeperAI) MarkMine(cell Cell) {
	ai.mines[cell] = true
	for _, sentence := range ai.knowledge {
		sentence.MarkMine(cell)
	}
}

func (ai *MinesweeperAI) MarkSafe(cell Cell) {
	ai.safes[cell] = true
	for _, sentence := range ai.knowledge {
		sentence.MarkSafe(cell)
	}
}

func (ai *MinesweeperAI) AppendKnowledge(cells []Cell, count int) {
	s := NewSentence(cells, count)
	ai.knowledge = append(ai.knowledge, &s)
}

func (ai *MinesweeperAI) RemoveKnowledge(sentence Sentence) {
	index := -1
	for i, originSentence := range ai.knowledge {
		if sentence.Equal(*originSentence) {
			index = i
			break
		}
	}

	if index >= 0 {
		ai.knowledge = append(ai.knowledge[:index], ai.knowledge[index+1:]...)
	}
}

func (ai *MinesweeperAI) IsInKnowledge(sentence Sentence) bool {
	for _, s := range ai.knowledge {
		if s.Equal(sentence) {
			return true
		}
	}
	return false
}

func (ai *MinesweeperAI) EqualKnowledge(other []*Sentence) bool {
	if len(ai.knowledge) != len(other) {
		return false
	}

	for _, aiSentence := range ai.knowledge {
		check := false
		for _, otherSentence := range other {
			if aiSentence.Equal(*otherSentence) {
				check = true
			}
		}
		if !check {
			return false
		}

	}
	return true
}

func (ai *MinesweeperAI) CopyKnowledge() []*Sentence {
	copy := make([]*Sentence, len(ai.knowledge))
	for i, sentence := range ai.knowledge {
		s := NewSentence(sentence.Cells, sentence.Count)
		copy[i] = &s
	}
	return copy
}

func (ai *MinesweeperAI) IsValidCell(i, j int) bool {
	return 0 <= i && i < ai.height && 0 <= j && j < ai.width
}

func (ai *MinesweeperAI) AddKnowledge(cell Cell, count int) {
	ai.movesMade[cell] = true

	ai.MarkSafe(cell)

	newSentenceCells := make([]Cell, 0)

	for i := cell.I - 1; i <= cell.I+1; i++ {
		for j := cell.J - 1; j <= cell.J+1; j++ {
			neighborCell := Cell{i, j}
			if cell.Equal(neighborCell) || !ai.IsValidCell(i, j) {
				continue
			}

			if ai.safes[neighborCell] {
				continue
			}

			if ai.mines[neighborCell] {
				count--
				continue
			}

			newSentenceCells = append(newSentenceCells, neighborCell)
		}
	}

	ai.AppendKnowledge(newSentenceCells, count)

	c := 0
	for {
		c++
		knowledgeCopy1 := ai.CopyKnowledge()

		for _, sentence := range knowledgeCopy1 {
			safe := sentence.KnownSafes()
			mine := sentence.KnownMines()

			if len(safe) > 0 {
				for _, cell := range safe {
					ai.safes[cell] = true
				}
				ai.RemoveKnowledge(*sentence)
			}

			if len(mine) > 0 {
				for _, cell := range mine {
					ai.mines[cell] = true
				}
				ai.RemoveKnowledge(*sentence)
			}
		}

		for safe := range ai.safes {
			ai.MarkSafe(safe)
		}

		for mine := range ai.mines {
			ai.MarkMine(mine)
		}

		knowledgeCopy2 := ai.CopyKnowledge()

		// knowledgeCopy2 é o conhecimento após todas as verificação acima
		// knowledgeCopy1 mantem o conhecimento inicial antes das verificações acima

		for _, sentenceKnowledge1 := range knowledgeCopy1 {
			for _, sentenceKnowledge2 := range knowledgeCopy2 {

				if sentenceKnowledge2.Equal(*sentenceKnowledge1) {
					continue
				}

				if !sentenceKnowledge1.IsSubsetCells(*sentenceKnowledge2) {
					continue
				}

				newCells := sentenceKnowledge2.DifferenceCells(*sentenceKnowledge1)
				if len(newCells) == 0 {
					continue
				}

				newSentence := NewSentence(newCells, sentenceKnowledge2.Count-sentenceKnowledge1.Count)

				if !ai.IsInKnowledge(newSentence) {
					ai.AppendKnowledge(newSentence.Cells, newSentence.Count)
				}
			}
		}

		if ai.EqualKnowledge(knowledgeCopy1) {
			break
		}

	}

}

func (ai *MinesweeperAI) MakeSafeMove() *Cell {
	for cell, value := range ai.safes {
		if value && !ai.movesMade[cell] && ai.IsValidCell(cell.I, cell.J) {
			return &cell
		}
	}
	return nil
}

func (ai *MinesweeperAI) MakeRandomMove() *Cell {
	for i := 0; i < ai.height; i++ {
		for j := 0; j < ai.width; j++ {
			move := Cell{i, j}
			if !ai.movesMade[move] && !ai.mines[move] && ai.IsValidCell(move.I, move.J) {
				return &move
			}
		}
	}
	return nil
}

func (ai *MinesweeperAI) MinesFound() map[Cell]bool {
	return ai.mines
}
