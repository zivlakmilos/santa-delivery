package game

import (
	"bytes"
	"image"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/zivlakmilos/santa-delivery/resources/images"

	_ "image/png"
)

type Gameplay struct {
	background *ebiten.Image
}

func NewGameplay() *Gameplay {
	bg, _, err := image.Decode(bytes.NewReader(images.BackgroundPng))
	if err != nil {
		log.Fatal(err)
	}

	return &Gameplay{
		background: ebiten.NewImageFromImage(bg),
	}
}

func (g *Gameplay) Update() (GameState, error) {
	return GameStateGame, nil
}

func (g *Gameplay) Draw(screen *ebiten.Image) {
	opt := &ebiten.DrawImageOptions{}
	screen.DrawImage(g.background, opt)
}
