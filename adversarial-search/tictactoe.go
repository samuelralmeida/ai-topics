package main

import (
	"math"
)

const X string = "X"
const O string = "O"

func initialState() [][]string {
	return [][]string{
		{"", "", ""},
		{"", "", ""},
		{"", "", ""},
	}
}

func player(board [][]string) string {
	moves := 0
	for _, row := range board {
		for _, cel := range row {
			if cel != "" {
				moves++
			}
		}
	}

	if moves%2 == 0 {
		return X
	}
	return O
}

func actions(board [][]string) [][]int {
	actions := [][]int{}
	for i, row := range board {
		for j, cel := range row {
			if cel == "" {
				actions = append(actions, []int{i, j})
			}
		}
	}
	return actions
}

func result(board [][]string, action []int) [][]string {
	result := initialState()

	for i, row := range board {
		copy(result[i], row)
	}

	result[action[0]][action[1]] = player(board)
	return result
}

func winner(board [][]string) string {

	// check rows
	if board[0][0] != "" && board[0][0] == board[0][1] && board[0][1] == board[0][2] {
		return board[0][0]
	} else if board[1][0] != "" && board[1][0] == board[1][1] && board[1][1] == board[1][2] {
		return board[1][0]
	} else if board[2][0] != "" && board[2][0] == board[2][1] && board[2][1] == board[2][2] {
		return board[2][0]
	}

	// check collumns
	if board[0][0] != "" && board[0][0] == board[1][0] && board[1][0] == board[2][0] {
		return board[0][0]
	} else if board[0][1] != "" && board[0][1] == board[1][1] && board[1][1] == board[2][1] {
		return board[0][1]
	} else if board[0][2] != "" && board[0][2] == board[1][2] && board[1][2] == board[2][2] {
		return board[0][2]
	}

	// check diagonals
	if board[0][0] != "" && board[0][0] == board[1][1] && board[1][1] == board[2][2] {
		return board[0][0]
	} else if board[0][2] != "" && board[0][2] == board[1][1] && board[1][1] == board[2][0] {
		return board[0][2]
	}

	return ""
}

func terminal(board [][]string) bool {
	if winner(board) != "" {
		return true
	}

	if len(actions(board)) == 0 {
		return true
	}

	return false
}

func utility(board [][]string) int {
	w := winner(board)
	if w == X {
		return 1
	} else if w == O {
		return -1
	} else {
		return 0
	}
}

func minimax(board [][]string) []int {
	if terminal(board) {
		return []int{}
	}

	player := player(board)
	actions := actions(board)

	bestAction := []int{}
	var bestScore *int = nil

	for _, action := range actions {
		if player == X {
			score := minValue(result(board, action), math.MinInt)
			if bestScore == nil || score > *bestScore {
				bestScore = &score
				bestAction = action
			}
		} else if player == O {
			score := maxValue(result(board, action), math.MaxInt)
			if bestScore == nil || score < *bestScore {
				bestScore = &score
				bestAction = action
			}
		}
	}

	return bestAction
}

func maxValue(board [][]string, highestCurrentValue int) int {
	if terminal(board) {
		return utility(board)
	}

	lowestPossibleValue := math.MinInt
	for _, action := range actions(board) {
		lowerValue := minValue(result(board, action), lowestPossibleValue)
		if lowerValue > highestCurrentValue {
			return lowerValue
		}
		lowestPossibleValue = max(lowestPossibleValue, lowerValue)
	}

	return lowestPossibleValue
}

func minValue(board [][]string, lowestCurrentValue int) int {
	if terminal(board) {
		return utility(board)
	}

	highestPossibleValue := math.MaxInt
	for _, action := range actions(board) {
		higherValue := maxValue(result(board, action), highestPossibleValue)
		if higherValue < lowestCurrentValue {
			return higherValue
		}
		highestPossibleValue = min(highestPossibleValue, higherValue)
	}

	return highestPossibleValue
}
