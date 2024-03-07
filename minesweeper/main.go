package main

import (
	"fmt"
	"html/template"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/samuelralmeida/ai-topics/minesweeper/minesweeper"
)

var (
	height, width, totalMines = 8, 8, 8
)

type TemplateRenderer struct {
	templates *template.Template
}

// Render renders a template document
func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {

	// Add global methods if data is a map
	if viewContext, isMap := data.(map[string]interface{}); isMap {
		viewContext["reverse"] = c.Echo().Reverse
	}

	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {

	g := minesweeper.NewGame(height, width, totalMines)
	ai := minesweeper.NewAI(height, width)

	reveled := make(map[minesweeper.Cell]bool)

	e := echo.New()

	renderer := &TemplateRenderer{
		templates: template.Must(template.ParseGlob("*.html")),
	}
	e.Renderer = renderer

	e.GET("/", func(c echo.Context) error {

		return c.Render(http.StatusOK, "board.html", map[string]interface{}{
			"name": "Dolly!",
		})

	})

	e.GET("/board", func(c echo.Context) error {
		resp := ""
		for i := 0; i < width; i++ {
			for j := 0; j < height; j++ {
				resp = resp + fmt.Sprintf(`<div class="board" id="board"><div class="cell" data-row="%d" data-col="%d"></div>`, i, j)
			}
		}
		return c.HTML(http.StatusOK, resp)
	})

	e.GET("/turn", func(c echo.Context) error {

		turn := game(g, ai, reveled)

		aiMessage := fmt.Sprintf(`
			<div class="ai-message">
				<p>%s</p>
			</div>
		`, turn.aiMessage)

		board := `<div class="board" id="board">`

		for i, row := range turn.board {
			for j, value := range row {
				board = board + fmt.Sprintf(`<div class="cell" data-row="%d" data-col="%d">%s</div>`, i, j, value)
			}
		}

		board = board + "</div>"

		gameMessage := fmt.Sprintf(`
			<div class="game-message">
				<p>%s</p>
			</div>
		`, turn.gameMessage)

		html := gameMessage + board + aiMessage

		return c.String(http.StatusOK, html)
	})

	e.Logger.Fatal(e.Start(":1324"))

}

type turn struct {
	aiMessage   string
	gameMessage string
	board       [][]string
	gameOver    bool
	move        minesweeper.Cell
}

func game(game *minesweeper.Minesweeper, ai *minesweeper.MinesweeperAI, reveled map[minesweeper.Cell]bool) turn {
	// height, width, totalMines := 8, 8, 8
	// game := minesweeper.NewGame(height, width, totalMines)
	// ai := minesweeper.NewAI(height, width)

	// reveled := make(map[minesweeper.Cell]bool)

	turn := turn{}

	move := ai.MakeSafeMove()
	if move == nil {
		move = ai.MakeRandomMove()
		if move == nil {
			game.GameOver = true
			turn.gameOver = true
			turn.gameMessage = "GAME: Não há mais movimentos para fazer."
		} else {
			turn.aiMessage = "IA: Não havia movimentos seguros par afazer, fiz uma escolha aleatória."
			turn.move = *move
		}
	} else {
		turn.aiMessage = "IA: Fiz um movimento seguro."
		turn.move = *move
	}

	if move != nil {
		if game.IsMine(*move) {
			game.Lost = true
			turn.gameOver = true
		} else {
			nearby := game.NearbyMines(*move)
			reveled[*move] = true
			ai.AddKnowledge(*move, nearby)
		}
	}

	if game.IsGameOver() {
		if game.Won(ai.MinesFound()) {
			turn.aiMessage = "IA: Eu sou foda!"
			turn.gameMessage = "IA venceu"
		} else {
			turn.aiMessage = "IA: Fui feito por um humano, culpe ele..."
			turn.gameMessage = "IA perdeu"
		}
	}

	board := make([][]string, width)
	for i := range board {
		board[i] = make([]string, height)
	}

	for i, row := range game.Board {
		for j := range row {
			cell := minesweeper.Cell{I: i, J: j}
			if game.IsMine(cell) && (game.Lost || game.GameOver) {
				board[i][j] = "X"
			} else if reveled[cell] {
				nearby := game.NearbyMines(cell)
				board[i][j] = fmt.Sprintf("%d", nearby)
			} else {
				board[i][j] = ""
			}

		}
	}

	turn.board = board
	return turn

	/*

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
	*/

}
