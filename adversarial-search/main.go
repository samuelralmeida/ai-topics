package main

import (
	"fmt"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

type Game struct {
	Board [][]string
}

func (g *Game) Update() error {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		// if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		mx, my := ebiten.CursorPosition()
		fmt.Println(mx, my)

		rowIdx := -1
		columnIdx := -1
		// square = -1

		if mx > 50 && mx < 150 && my > 50 && my < 150 {
			rowIdx = 0
			columnIdx = 0
		} else if mx > 50 && mx < 150 && my > 150 && my < 250 {
			rowIdx = 1
			columnIdx = 0
		} else if mx > 50 && mx < 150 && my > 250 && my < 350 {
			rowIdx = 2
			columnIdx = 0
		} else if mx > 150 && mx < 250 && my > 50 && my < 150 {
			rowIdx = 0
			columnIdx = 1
		} else if mx > 150 && mx < 250 && my > 150 && my < 250 {
			rowIdx = 1
			columnIdx = 1
		} else if mx > 150 && mx < 250 && my > 250 && my < 350 {
			rowIdx = 2
			columnIdx = 1
		} else if mx > 250 && mx < 350 && my > 50 && my < 150 {
			rowIdx = 0
			columnIdx = 2
		} else if mx > 250 && mx < 350 && my > 150 && my < 250 {
			rowIdx = 1
			columnIdx = 2
		} else if mx > 250 && mx < 350 && my > 250 && my < 350 {
			rowIdx = 2
			columnIdx = 2
		}

		if rowIdx == -1 || columnIdx == -1 {
			return nil
		}

		fmt.Println(g.Board)

		g.Board[rowIdx][columnIdx] = "O"

		fmt.Println(g.Board)

		// if mx > 50 && mx < 150 {
		// 	rowIdx = 0
		// } else if mx > 150 && mx < 250 {
		// 	rowIdx = 1
		// } else if mx > 250 && mx < 350 {
		// 	rowIdx = 2
		// }

		// if my > 50 && my < 150 {
		// 	columnIdx = 0
		// } else if my > 150 && my < 250 {
		// 	columnIdx = 1
		// } else if my > 250 && my < 350 {
		// 	columnIdx = 2
		// }

	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	vector.StrokeLine(screen, 50, 150, 350, 150, 3, color.White, false)
	vector.StrokeLine(screen, 50, 250, 350, 250, 3, color.White, false)
	vector.StrokeLine(screen, 150, 50, 150, 350, 3, color.White, false)
	vector.StrokeLine(screen, 250, 50, 250, 350, 3, color.White, false)

	tt, err := opentype.Parse(fonts.MPlus1pRegular_ttf)
	if err != nil {
		log.Fatal(err)
	}
	bigText, err := opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    50,
		DPI:     72,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatal(err)
	}

	correction := map[int]int{0: 50, 1: 150, 2: 250}
	for i, row := range g.Board {
		for j, collumn := range row {
			if collumn != "" {
				x := correction[j]
				y := correction[i]
				text.Draw(screen, collumn, bigText, x+30, y+70, color.White)
			}
		}
	}

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 400, 400
}

func main() {
	ebiten.SetWindowSize(400, 400)
	ebiten.SetWindowTitle("Play Tic-Tac-Toe")
	game := &Game{
		Board: [][]string{
			{"", "", ""},
			{"", "", ""},
			{"", "", ""},
		},
	}
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}

// func drawBoard(screen *ebiten.Image) {
// 	// tileOrigin := ebiten.Vector{
// 	// 	X: float64(screenWidth/2 - (1.5*tileSize)),
// 	// 	Y: float64(screenHeight/2 - (1.5*tileSize)),
// 	// }
// }
