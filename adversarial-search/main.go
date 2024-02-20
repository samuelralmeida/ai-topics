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
	Board           [][]string
	BackgroundBoard [][]RectangleArea
	Measurements
	User            string
	Player          string
	UserStartButton RectangleArea
	AiStartButton   RectangleArea
	Title           string
	PlayAgainButton RectangleArea
	GameOver        bool
}

func (g *Game) Update() error {

	// menu
	if g.User == "" {
		if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
			mx, my := ebiten.CursorPosition()

			if g.UserStartButton.IsPointInside(mx, my) {
				g.User = "X"
			} else if g.AiStartButton.IsPointInside(mx, my) {
				g.User = "O"
			}
		}
		return nil
	}

	gameOver := terminal(g.Board)
	player := player(g.Board)

	if gameOver {
		g.GameOver = true
		winner := winner(g.Board)
		if winner == "" {
			g.Title = "Game over: Tie."
		} else {
			g.Title = fmt.Sprintf("Game over: %s wins.", winner)
		}

		if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
			mx, my := ebiten.CursorPosition()
			if g.PlayAgainButton.IsPointInside(mx, my) {
				g.Board = initialState()
				g.GameOver = false
				g.User = ""
			}
		}

		return nil
	}

	// user
	if g.User == player {
		g.Title = fmt.Sprintf("Play as %s", g.User)
		if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
			mx, my := ebiten.CursorPosition()

			rowIdx := -1
			columnIdx := -1

			for i, row := range g.BackgroundBoard {
				for j, area := range row {
					if area.IsPointInside(mx, my) {
						rowIdx = i
						columnIdx = j
						break
					}
				}
			}

			if rowIdx == -1 || columnIdx == -1 {
				return nil
			}

			if g.Board[rowIdx][columnIdx] == "" {
				g.Board[rowIdx][columnIdx] = player
			}
		}
		return nil
	}

	// ai
	g.Title = "AI thinking..."
	move := minimax(g.Board)
	board := result(g.Board, move)
	g.Board = board
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

	if g.User == "" {
		userButtonLines := g.UserStartButton.Lines()
		for _, l := range userButtonLines {
			vector.StrokeLine(screen, float32(l.X0), float32(l.Y0), float32(l.X1), float32(l.Y1), 3, color.White, false)
		}

		msg := "I START"
		textWidth := font.MeasureString(bigText, msg).Ceil()
		textX := (g.Width - textWidth) / 2
		text.Draw(screen, msg, bigText, textX, g.Margin()+g.Space()+20, color.White)

		aiButtonLines := g.AiStartButton.Lines()
		for _, l := range aiButtonLines {
			vector.StrokeLine(screen, float32(l.X0), float32(l.Y0), float32(l.X1), float32(l.Y1), 3, color.White, false)
		}

		msg = "AI STARTS"
		textWidth = font.MeasureString(bigText, msg).Ceil()
		textX = (g.Width - textWidth) / 2
		text.Draw(screen, msg, bigText, textX, g.Margin()+g.Space()*4+20, color.White)

		return
	}

	if g.GameOver {
		// playAgainButton := g.PlayAgainButton.Lines()
		// for _, l := range playAgainButton {
		// 	vector.StrokeLine(screen, float32(l.X0), float32(l.Y0), float32(l.X1), float32(l.Y1), 3, color.White, false)
		// }

		msg := "PLAY AGAIN"
		textWidth := font.MeasureString(bigText, msg).Ceil()
		textX := (g.Width - textWidth) / 2
		text.Draw(screen, msg, bigText, textX, g.Margin()+g.Space()*7+20, color.White)
	}

	text.Draw(screen, g.Title, bigText, g.Space(), g.Space()*2, color.White)

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

	measurements := Measurements{
		Width:  width,
		Height: height,
	}

	userStartButton := NewRectangleArea(measurements.Space()*6, measurements.Space()*2, measurements.Margin(), measurements.Margin())
	aiStartButton := NewRectangleArea(measurements.Space()*6, measurements.Space()*2, measurements.Margin(), measurements.Margin()+measurements.Space()*3)
	playAgainButton := NewRectangleArea(measurements.Space()*6, measurements.Space()*2, measurements.Margin(), measurements.Margin()+measurements.Space()*6)

	game := &Game{
		Board: [][]string{
			{"", "", ""},
			{"", "", ""},
			{"", "", ""},
		},
		Measurements:    measurements,
		User:            "",
		Player:          "X",
		UserStartButton: userStartButton,
		AiStartButton:   aiStartButton,
		PlayAgainButton: playAgainButton,
		BackgroundBoard: NewBackgroundBoard(measurements.Square(), measurements.Square(), measurements.Margin(), measurements.Margin()),
	}
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}

type RectangleArea struct {
	Width  int
	Height int
	X      int
	Y      int
}

type Line struct {
	X0 int
	Y0 int
	X1 int
	Y1 int
}

func (ra *RectangleArea) TopHorizontalLine() Line {
	return Line{
		X0: ra.X,
		Y0: ra.Y,
		X1: ra.X + ra.Width,
		Y1: ra.Y,
	}
}

func (ra *RectangleArea) BottomHorizontalLine() Line {
	return Line{
		X0: ra.X,
		Y0: ra.Y + ra.Height,
		X1: ra.X + ra.Width,
		Y1: ra.Y + ra.Height,
	}
}

func (ra *RectangleArea) LeftVerticalLine() Line {
	return Line{
		X0: ra.X,
		Y0: ra.Y,
		X1: ra.X,
		Y1: ra.Y + ra.Height,
	}
}

func (ra *RectangleArea) RightVerticalLine() Line {
	return Line{
		X0: ra.X + ra.Width,
		Y0: ra.Y,
		X1: ra.X + ra.Width,
		Y1: ra.Y + ra.Height,
	}
}

func NewRectangleArea(width, height, relativeX, relativeY int) RectangleArea {
	return RectangleArea{Width: width, Height: height, X: relativeX, Y: relativeY}
}

func (ra *RectangleArea) IsPointInside(x, y int) bool {
	return x > ra.X && x < (ra.X+ra.Width) && y > ra.Y && y < (ra.Y+ra.Height)
}

func (ra *RectangleArea) Lines() []Line {
	return []Line{
		ra.TopHorizontalLine(),
		ra.BottomHorizontalLine(),
		ra.LeftVerticalLine(),
		ra.RightVerticalLine(),
	}
}

func NewBackgroundBoard(width, height, relativeX, relativeY int) [][]RectangleArea {

	area00 := NewRectangleArea(width, height, relativeX, relativeY)
	area01 := NewRectangleArea(width, height, relativeX+width, relativeY)
	area02 := NewRectangleArea(width, height, relativeX+width*2, relativeY)
	area10 := NewRectangleArea(width, height, relativeX, relativeY+height)
	area11 := NewRectangleArea(width, height, relativeX+width, relativeY+height)
	area12 := NewRectangleArea(width, height, relativeX+width*2, relativeY+height)
	area20 := NewRectangleArea(width, height, relativeX, relativeY+height*2)
	area21 := NewRectangleArea(width, height, relativeX+width, relativeY+height*2)
	area22 := NewRectangleArea(width, height, relativeX+width*2, relativeY+height*2)

	return [][]RectangleArea{
		{area00, area01, area02},
		{area10, area11, area12},
		{area20, area21, area22},
	}
}
