package game

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type GameState int

const (
	GameStateMenu GameState = iota
	GameStateGame
	GameStateOver
	GameStateCount
)

type Scene interface {
	Update() (GameState, error)
	Draw(screen *ebiten.Image)
}

type Game struct {
	states [GameStateCount]Scene
	state  GameState
}

func NewGame() *Game {
	g := &Game{
		state: GameStateMenu,
	}

	return g
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0x00, 0x00, 0x00, 0xFF})
	ebitenutil.DebugPrint(screen, "Hello World")
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func StartGame() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Santa's Delivery")
	game := NewGame()
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
