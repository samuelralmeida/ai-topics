package main

import (
	"fmt"
	"time"

	"github.com/samuelralmeida/ai-topics/minesweeper/minesweeper"
)

func main() {
	height, width, totalMines := 8, 8, 8
	game := minesweeper.NewGame(height, width, totalMines)
	ai := minesweeper.NewAI(height, width)

	reveled := make(map[minesweeper.Cell]bool)

	for {

		// print board
		for i, row := range game.Board {
			for j := range row {
				cell := minesweeper.Cell{I: i, J: j}
				if game.IsMine(cell) && (game.Lost || game.GameOver) {
					fmt.Print("|X")
				} else if reveled[cell] {
					nearby := game.NearbyMines(cell)
					fmt.Printf("|%d", nearby)
				} else {
					fmt.Print("| ")
				}

			}
			fmt.Println("|")
		}

		// check end game
		if game.IsGameOver() {
			if game.Won(ai.MinesFound()) {
				fmt.Println("You win")
			} else {
				fmt.Println("You lose")
			}

			break
		}

		//
		move := ai.MakeSafeMove()
		if move == nil {
			move = ai.MakeRandomMove()
			if move == nil {
				game.GameOver = true
				fmt.Println("No moves left to make.")
			} else {
				fmt.Println("No known safe moves, AI making random move.", *move)
			}
		} else {
			fmt.Println("AI making safe move.", *move)
		}

		if game.GameOver {
			continue
		}

		if move != nil {
			if game.IsMine(*move) {
				game.Lost = true
			} else {
				nearby := game.NearbyMines(*move)
				reveled[*move] = true
				ai.AddKnowledge(*move, nearby)
			}
		}

		time.Sleep(time.Millisecond * 300)

	}

}
