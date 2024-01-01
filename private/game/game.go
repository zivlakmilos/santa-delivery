package game

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

type GameState int

const (
	GameStateMenu GameState = iota
	GameStateGame
	GameStateOver
	GameStateCount
)

const (
	screenWidth  = 720
	screenHeight = 480
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

	g.states[GameStateMenu] = NewMenu()
	g.states[GameStateGame] = NewGameplay()

	return g
}

func (g *Game) Update() error {
	state := g.states[g.state]
	nextState, err := state.Update()
	if err != nil {
		return err
	}

	g.state = nextState
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	state := g.states[g.state]

	screen.Clear()
	state.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func StartGame() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Santa's Delivery")
	game := NewGame()
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
