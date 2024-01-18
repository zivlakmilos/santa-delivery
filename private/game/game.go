package game

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/zivlakmilos/santa-delivery/private/constant"
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
	return constant.ScreenWidth, constant.ScreenHeight
}

func StartGame() {
	ebiten.SetWindowSize(constant.ScreenWidth, constant.ScreenHeight)
	ebiten.SetWindowTitle("Santa's Delivery")
	game := NewGame()
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
