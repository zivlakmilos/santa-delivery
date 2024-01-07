package game

import (
	"bytes"
	"image"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/zivlakmilos/santa-delivery/private/object"
	"github.com/zivlakmilos/santa-delivery/resources/images"

	_ "image/png"
)

type Gameplay struct {
	background *ebiten.Image
	camera     *object.Camera
	player     *object.Player
}

func NewGameplay() *Gameplay {
	bg, _, err := image.Decode(bytes.NewReader(images.BackgroundPng))
	if err != nil {
		log.Fatal(err)
	}

	return &Gameplay{
		background: ebiten.NewImageFromImage(bg),
		camera:     object.NewCamera(0, 0, screenWidth, screenHeight),
		player:     object.NewPlayer(100, 100),
	}
}

func (g *Gameplay) Update() (GameState, error) {
	err := g.player.Update()
	if err != nil {
		return GameStateGame, err
	}

	g.camera.MoveTo(g.player.GetPos())

	return GameStateGame, nil
}

func (g *Gameplay) Draw(screen *ebiten.Image) {
	opt := &ebiten.DrawImageOptions{}
	screen.DrawImage(g.background, opt)

	bg, _, err := image.Decode(bytes.NewReader(images.Tile1Png))
	if err != nil {
		return
	}
	opt = &ebiten.DrawImageOptions{}
	opt.GeoM.Translate(10.0-g.camera.GetPos().X, 10.0-g.camera.GetPos().Y)
	opt.GeoM.Scale(0.5, 0.5)
	img := ebiten.NewImageFromImage(bg)
	screen.DrawImage(img, opt)

	g.player.Draw(screen, g.camera)
}
