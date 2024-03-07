package minesweeper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSentence_MarkMine(t *testing.T) {
	sentence := NewSentence([]Cell{{1, 0}, {1, 1}, {1, 2}}, 2)

	assert.Equal(t, sentence.Count, 2)
	assert.Equal(t, len(sentence.Cells), 3)

	sentence.MarkMine(Cell{1, 1})

	assert.Equal(t, sentence.Count, 1)
	assert.Equal(t, len(sentence.Cells), 2)
	assert.Equal(t, Cell{1, 0}, sentence.Cells[0])
	assert.Equal(t, Cell{1, 2}, sentence.Cells[1])
}

func TestSentence_MarkSafe(t *testing.T) {
	sentence := NewSentence([]Cell{{1, 0}, {1, 1}, {1, 2}}, 2)

	assert.Equal(t, sentence.Count, 2)
	assert.Equal(t, len(sentence.Cells), 3)

	sentence.MarkSafe(Cell{1, 1})

	assert.Equal(t, sentence.Count, 2)
	assert.Equal(t, len(sentence.Cells), 2)
	assert.Equal(t, Cell{1, 0}, sentence.Cells[0])
	assert.Equal(t, Cell{1, 2}, sentence.Cells[1])
}

func TestSentence_Equal(t *testing.T) {
	sentence1 := NewSentence([]Cell{{1, 0}, {1, 1}, {1, 2}}, 2)

	sentence2 := NewSentence([]Cell{{1, 0}, {1, 1}}, 2)
	assert.False(t, sentence1.Equal(sentence2))

	sentence3 := NewSentence([]Cell{{1, 0}, {1, 1}, {1, 2}}, 3)
	assert.False(t, sentence1.Equal(sentence3))

	sentence4 := NewSentence([]Cell{{1, 0}, {1, 1}, {1, 3}}, 2)
	assert.False(t, sentence1.Equal(sentence4))

	sentence5 := NewSentence([]Cell{{1, 0}, {1, 1}, {1, 2}}, 2)
	assert.True(t, sentence1.Equal(sentence5))
}

func TestSentence_DifferenceCells(t *testing.T) {
	sentence1 := NewSentence([]Cell{{1, 0}, {1, 1}, {1, 2}}, 2)
	sentence2 := NewSentence([]Cell{{2, 0}, {2, 1}, {1, 0}}, 2)

	cells := sentence2.DifferenceCells(sentence1)

	assert.Len(t, cells, 2)
	assert.Equal(t, Cell{2, 0}, cells[0])
	assert.Equal(t, Cell{2, 1}, cells[1])
}

func TestSentence_IsSubsetCells(t *testing.T) {
	sentence1 := NewSentence([]Cell{{1, 0}, {1, 1}, {1, 2}}, 2)
	sentence2 := NewSentence([]Cell{{1, 0}, {1, 1}}, 2)

	assert.True(t, sentence2.IsSubsetCells(sentence1))

	sentence3 := NewSentence([]Cell{{1, 0}, {2, 1}}, 2)
	assert.False(t, sentence3.IsSubsetCells(sentence1))
}

//

func TestMinesweeperAI_MarkMine(t *testing.T) {
	ai := NewAI(8, 8)

	s := NewSentence([]Cell{{1, 0}, {1, 1}, {1, 2}}, 2)

	ai.knowledge = []*Sentence{&s}

	ai.MarkMine(Cell{1, 1})

	sentence := ai.knowledge[0]

	assert.Equal(t, sentence.Count, 1)
	assert.Equal(t, len(sentence.Cells), 2)
	assert.Equal(t, Cell{1, 0}, sentence.Cells[0])
	assert.Equal(t, Cell{1, 2}, sentence.Cells[1])
}

func TestMinesweeperAI_MarkSafe(t *testing.T) {
	ai := NewAI(8, 8)

	s := NewSentence([]Cell{{1, 0}, {1, 1}, {1, 2}}, 2)

	ai.knowledge = []*Sentence{&s}

	ai.MarkSafe(Cell{1, 1})

	sentence := ai.knowledge[0]

	assert.Equal(t, sentence.Count, 2)
	assert.Equal(t, len(sentence.Cells), 2)
	assert.Equal(t, Cell{1, 0}, sentence.Cells[0])
	assert.Equal(t, Cell{1, 2}, sentence.Cells[1])
}

func TestMinesweeperAI_AppendKnowledge(t *testing.T) {
	ai := NewAI(8, 8)

	ai.AppendKnowledge([]Cell{{1, 0}, {1, 1}, {1, 2}}, 2)

	sentence := ai.knowledge[0]

	assert.Equal(t, sentence.Count, 2)
	assert.Equal(t, len(sentence.Cells), 3)
}

func TestMinesweeperAI_CopyKnowledge(t *testing.T) {
	ai := NewAI(8, 8)

	s := NewSentence([]Cell{{1, 0}, {1, 1}, {1, 2}}, 2)

	ai.knowledge = []*Sentence{&s}

	copy := ai.CopyKnowledge()

	assert.True(t, ai.knowledge[0].Equal(*copy[0]))

	ai.MarkSafe(Cell{1, 1})

	assert.False(t, ai.knowledge[0].Equal(*copy[0]))
}

func TestMinesweeperAI_IsValidCell(t *testing.T) {
	ai := NewAI(8, 8)

	assert.False(t, ai.IsValidCell(-1, 0))
	assert.False(t, ai.IsValidCell(0, -1))
	assert.False(t, ai.IsValidCell(8, -1))
	assert.False(t, ai.IsValidCell(0, 8))
	assert.True(t, ai.IsValidCell(0, 0))
	assert.True(t, ai.IsValidCell(7, 7))
}

func TestMinesweeperAI_RemoveKnowledge(t *testing.T) {
	ai := NewAI(8, 8)

	s1 := NewSentence([]Cell{{1, 0}, {1, 1}, {1, 2}}, 2)
	s2 := NewSentence([]Cell{{1, 3}, {1, 4}, {1, 5}}, 1)
	s3 := NewSentence([]Cell{{1, 6}}, 1)

	ai.knowledge = []*Sentence{&s1, &s2, &s3}

	ai.RemoveKnowledge(NewSentence([]Cell{{1, 3}, {1, 4}, {1, 5}}, 1))

	assert.Len(t, ai.knowledge, 2)
}

func TestMinesweeperAI_EqualKnowledge(t *testing.T) {
	ai := NewAI(8, 8)

	s1 := NewSentence([]Cell{{1, 0}, {1, 1}, {1, 2}}, 2)
	s2 := NewSentence([]Cell{{1, 3}, {1, 4}, {1, 5}}, 1)
	s3 := NewSentence([]Cell{{1, 6}}, 1)

	ai.knowledge = []*Sentence{&s1, &s2, &s3}
	otherKnowlegde1 := []*Sentence{&s1, &s2, &s3}

	assert.True(t, ai.EqualKnowledge(otherKnowlegde1))

	s4 := NewSentence([]Cell{{4, 6}}, 1)
	otherKnowlegde2 := []*Sentence{&s1, &s3, &s4}

	assert.False(t, ai.EqualKnowledge(otherKnowlegde2))
}

func TestMinesweeperAI_AddKnowledge(t *testing.T) {
	ai := NewAI(8, 8)

	ai.AddKnowledge(Cell{0, 0}, 1)
	ai.AddKnowledge(Cell{0, 1}, 2)
	ai.AddKnowledge(Cell{0, 2}, 2)
	ai.AddKnowledge(Cell{0, 3}, 2)
	ai.AddKnowledge(Cell{0, 4}, 1)
	ai.AddKnowledge(Cell{0, 5}, 1)
	ai.AddKnowledge(Cell{0, 6}, 0)
	ai.AddKnowledge(Cell{1, 7}, 0)
	ai.AddKnowledge(Cell{2, 7}, 0)
	ai.AddKnowledge(Cell{0, 7}, 0)
	ai.AddKnowledge(Cell{1, 6}, 0)
	ai.AddKnowledge(Cell{1, 5}, 1)
	ai.AddKnowledge(Cell{3, 6}, 0)
	ai.AddKnowledge(Cell{1, 3}, 3)

	assert.Len(t, ai.movesMade, 14)
}
