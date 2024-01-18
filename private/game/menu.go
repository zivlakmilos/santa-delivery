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
	"github.com/zivlakmilos/santa-delivery/private/constant"
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
	return m.handleInput(), nil
}

func (m *Menu) Draw(screen *ebiten.Image) {
	opt := &ebiten.DrawImageOptions{}
	screen.DrawImage(m.background, opt)

	text.Draw(screen, "Press SPACE to Start", m.font, constant.ScreenWidth/2-125, constant.ScreenHeight/2+175, color.White)
	text.Draw(screen, "Pres ESC to Quit", m.font, constant.ScreenWidth/2-100, constant.ScreenHeight/2+225, color.White)
}

func (m *Menu) handleInput() GameState {
	if ebiten.IsKeyPressed(ebiten.KeyQ) || ebiten.IsKeyPressed(ebiten.KeyEscape) {
		os.Exit(0)
	}

	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		return GameStateGame
	}

	return GameStateMenu
}
