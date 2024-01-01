package game

import (
	"bytes"
	"image"
	"image/color"
	"log"
	"os"

	"github.com/golang/freetype/truetype"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/zivlakmilos/santa-delivery/resources/images"
	"golang.org/x/image/font"
	"golang.org/x/image/font/gofont/goregular"

	_ "image/png"
)

type Menu struct {
	background *ebiten.Image
	font       font.Face
}

func NewMenu() *Menu {
	bg, _, err := image.Decode(bytes.NewReader(images.BackgroundPng))
	if err != nil {
		log.Fatal(err)
	}

	font, err := truetype.Parse(goregular.TTF)
	if err != nil {
		log.Fatal(err)
	}

	return &Menu{
		background: ebiten.NewImageFromImage(bg),
		font:       truetype.NewFace(font, &truetype.Options{Size: 25}),
	}
}

func (m *Menu) Update() (GameState, error) {
	m.handleInput()

	return GameStateMenu, nil
}

func (m *Menu) Draw(screen *ebiten.Image) {
	opt := &ebiten.DrawImageOptions{}
	screen.DrawImage(m.background, opt)

	text.Draw(screen, "Press SPACE to Start", m.font, screenWidth/2-125, screenHeight/2+175, color.White)
	text.Draw(screen, "Pres ESC to Quit", m.font, screenWidth/2-100, screenHeight/2+225, color.White)
}

func (m *Menu) handleInput() {
	if ebiten.IsKeyPressed(ebiten.KeyQ) || ebiten.IsKeyPressed(ebiten.KeyEscape) {
		os.Exit(0)
	}
}
