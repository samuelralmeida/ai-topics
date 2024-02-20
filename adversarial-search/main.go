package main

import (
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

var bigText font.Face

func init() {
	tt, err := opentype.Parse(fonts.MPlus1pRegular_ttf)
	if err != nil {
		log.Fatal(err)
	}
	bigText, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    50,
		DPI:     72,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatal(err)
	}
}

type Measurements struct {
	Width  int
	Height int
}

func (m *Measurements) Space() int {
	return m.Width / 12
}

func (m *Measurements) Margin() int {
	return m.Space() * 3
}

func (m *Measurements) Square() int {
	return m.Space() * 2
}

type Game struct {
	Board [][]string
	Measurements
	User   string
	Player string
}

func (g *Game) Update() error {

	if g.User == "" {
		if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
			mx, my := ebiten.CursorPosition()

			if mx > g.Margin() && mx < g.Margin()+g.Space()*6 && my > g.Margin() && my < g.Margin()+g.Space()*2 {
				g.User = "X"
			} else if mx > g.Margin() && mx < g.Margin()+g.Space()*6 && my > g.Margin()+g.Space()*3 && my < g.Margin()+g.Space()*5 {
				g.User = "O"
			}
		}
		return nil
	}

	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		mx, my := ebiten.CursorPosition()

		rowIdx := -1
		columnIdx := -1

		if mx > g.Margin() && mx < g.Margin()+g.Square() && my > g.Margin() && my < g.Margin()+g.Square() {
			rowIdx = 0
			columnIdx = 0
		} else if mx > g.Margin() && mx < g.Margin()+g.Square() && my > g.Margin()+g.Square() && my < g.Margin()+g.Square()*2 {
			rowIdx = 1
			columnIdx = 0
		} else if mx > g.Margin() && mx < g.Margin()+g.Square() && my > g.Margin()+g.Square()*2 && my < g.Margin()+g.Square()*3 {
			rowIdx = 2
			columnIdx = 0
		} else if mx > g.Margin()+g.Square() && mx < g.Margin()+g.Square()*2 && my > g.Margin() && my < g.Margin()+g.Square() {
			rowIdx = 0
			columnIdx = 1
		} else if mx > g.Margin()+g.Square() && mx < g.Margin()+g.Square()*2 && my > g.Margin()+g.Square() && my < g.Margin()+g.Square()*2 {
			rowIdx = 1
			columnIdx = 1
		} else if mx > g.Margin()+g.Square() && mx < g.Margin()+g.Square()*2 && my > g.Margin()+g.Square()*2 && my < g.Margin()+g.Square()*3 {
			rowIdx = 2
			columnIdx = 1
		} else if mx > g.Margin()+g.Square()*2 && mx < g.Margin()+g.Square()*3 && my > g.Margin() && my < g.Margin()+g.Square() {
			rowIdx = 0
			columnIdx = 2
		} else if mx > g.Margin()+g.Square()*2 && mx < g.Margin()+g.Square()*3 && my > g.Margin()+g.Square() && my < g.Margin()+g.Square()*2 {
			rowIdx = 1
			columnIdx = 2
		} else if mx > g.Margin()+g.Square()*2 && mx < g.Margin()+g.Square()*3 && my > g.Margin()+g.Square()*2 && my < g.Margin()+g.Square()*3 {
			rowIdx = 2
			columnIdx = 2
		}

		if rowIdx == -1 || columnIdx == -1 {
			return nil
		}

		g.Board[rowIdx][columnIdx] = g.Player
		g.ChangePlayer()

	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

	// var h float32 = 0
	// for h <= float32(g.Measurements.Width) {
	// 	h = h + 50
	// 	vector.StrokeLine(screen, 0, h, float32(g.Measurements.Width), h, 3, color.RGBA{255, 0, 0, 0}, false)
	// 	vector.StrokeLine(screen, h, 0, h, float32(g.Measurements.Width), 3, color.RGBA{255, 0, 0, 0}, false)
	// }

	if g.User == "" {
		vector.StrokeLine(screen, float32(g.Margin()), float32(g.Margin()), float32(g.Margin()+g.Space()*6), float32(g.Margin()), 3, color.White, false)
		vector.StrokeLine(screen, float32(g.Margin()), float32(g.Margin()), float32(g.Margin()), float32(g.Margin()+g.Space()*2), 3, color.White, false)
		vector.StrokeLine(screen, float32(g.Margin()), float32(g.Margin()+g.Space()*2), float32(g.Margin()+g.Space()*6), float32(g.Margin()+g.Space()*2), 3, color.White, false)
		vector.StrokeLine(screen, float32(g.Margin()+g.Space()*6), float32(g.Margin()), float32(g.Margin()+g.Space()*6), float32(g.Margin()+g.Space()*2), 3, color.White, false)

		msg := "I START"
		textWidth := font.MeasureString(bigText, msg).Ceil()
		textX := (g.Width - textWidth) / 2
		text.Draw(screen, msg, bigText, textX, g.Margin()+g.Space()+20, color.White)

		vector.StrokeLine(screen, float32(g.Margin()), float32(g.Margin()+g.Space()*3), float32(g.Margin()+g.Space()*6), float32(g.Margin()+g.Space()*3), 3, color.White, false)
		vector.StrokeLine(screen, float32(g.Margin()), float32(g.Margin()+g.Space()*5), float32(g.Margin()+g.Space()*6), float32(g.Margin()+g.Space()*5), 3, color.White, false)
		vector.StrokeLine(screen, float32(g.Margin()), float32(g.Margin()+g.Space()*3), float32(g.Margin()), float32(g.Margin()+g.Space()*5), 3, color.White, false)
		vector.StrokeLine(screen, float32(g.Margin()+g.Space()*6), float32(g.Margin()+g.Space()*3), float32(g.Margin()+g.Space()*6), float32(g.Margin()+g.Space()*5), 3, color.White, false)

		msg = "AI STARTS"
		textWidth = font.MeasureString(bigText, msg).Ceil()
		textX = (g.Width - textWidth) / 2
		text.Draw(screen, msg, bigText, textX, g.Margin()+g.Space()*4+20, color.White)

		return
	}

	instruction := "Your turn"
	if g.User != g.Player {
		instruction = "Computer thinking..."
	}

	text.Draw(screen, instruction, bigText, g.Space(), g.Space()*2, color.White)

	vector.StrokeLine(screen, float32(g.Margin()), float32(g.Margin()+g.Square()), float32(g.Margin()+g.Square()*3), float32(g.Measurements.Margin()+g.Measurements.Square()), 3, color.White, false)
	vector.StrokeLine(screen, float32(g.Margin()), float32(g.Margin()+g.Square()*2), float32(g.Margin()+g.Square()*3), float32(g.Margin()+g.Square()*2), 3, color.White, false)
	vector.StrokeLine(screen, float32(g.Margin()+g.Square()), float32(g.Margin()), float32(g.Margin()+g.Square()), float32(g.Margin()+g.Square()*3), 3, color.White, false)
	vector.StrokeLine(screen, float32(g.Margin()+g.Square()*2), float32(g.Margin()), float32(g.Margin()+g.Square()*2), float32(g.Margin()+g.Square()*3), 3, color.White, false)

	correction := map[int]int{0: g.Margin(), 1: g.Margin() + g.Square(), 2: g.Margin() + g.Square()*2}
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
	return g.Measurements.Width, g.Measurements.Height
}

func (g *Game) ChangePlayer() {
	if g.Player == "X" {
		g.Player = "O"
		return
	}
	g.Player = "X"
}

func main() {
	width := 600
	height := 600

	ebiten.SetWindowSize(width, height)
	ebiten.SetWindowTitle("Play Tic-Tac-Toe")

	game := &Game{
		Board: [][]string{
			{"", "", ""},
			{"", "", ""},
			{"", "", ""},
		},
		Measurements: Measurements{
			Width:  width,
			Height: height,
		},
		User:   "",
		Player: "X",
	}
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
