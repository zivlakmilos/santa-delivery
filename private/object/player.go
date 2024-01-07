package object

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type Player struct {
	pos Vector
}

func NewPlayer(x, y float64) *Player {
	return &Player{
		pos: NewVector(x, y),
	}
}

func (p *Player) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		p.pos.X += 10
	}
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		p.pos.X -= 10
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		p.pos.Y += 10
	}
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		p.pos.Y -= 10
	}

	return nil
}

func (p *Player) Draw(screen *ebiten.Image, camera *Camera) {
	img := ebiten.NewImage(25, 25)
	img.Fill(color.White)

	opt := &ebiten.DrawImageOptions{}
	opt.GeoM.Translate(p.pos.X-camera.pos.X+camera.width/2, p.pos.Y-camera.pos.Y+camera.height/2)
	screen.DrawImage(img, opt)
}

func (p *Player) GetPos() Vector {
	return p.pos
}
