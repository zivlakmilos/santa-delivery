package object

import (
	"bytes"
	"image"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/zivlakmilos/santa-delivery/private/utils"
	"github.com/zivlakmilos/santa-delivery/resources/images"
)

type PlayerState int

const (
	PlayerStateIdle PlayerState = iota
	PlayerStateWalk
	PlayerStateCount // NOTE: Must be last
)

type Player struct {
	pos   Vector
	vel   Vector
	face  Vector
	state PlayerState

	sprites         [PlayerStateCount][]*ebiten.Image
	animationIdx    int
	animationFrames int
}

func NewPlayer(x, y float64) *Player {
	p := &Player{
		pos:   NewVector(x, y),
		vel:   NewVector(0, 0),
		state: PlayerStateIdle,
	}

	p.loadSprites()

	return p
}

func (p *Player) Update() error {
	p.vel.X = 0
	p.vel.Y = 0

	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		p.vel.X += 5
	}
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		p.vel.X -= 5
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		p.vel.Y += 5
	}
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		p.vel.Y -= 5
	}
	p.pos.Add(p.vel)

	if p.vel.X != 0 {
		p.face.X = p.vel.X
		p.setState(PlayerStateWalk)
	} else {
		p.setState(PlayerStateIdle)
	}

	p.updateAnimation()

	return nil
}

func (p *Player) Draw(screen *ebiten.Image, camera *Camera) {
	img := p.sprites[p.state][p.animationIdx]
	if p.face.X < 0 {
		img = utils.FlipImageHorizontal(img)
	}

	opt := &ebiten.DrawImageOptions{}
	opt.GeoM.Translate(p.pos.X-camera.pos.X+camera.width/2, p.pos.Y-camera.pos.Y+camera.height/2)

	screen.DrawImage(img, opt)
}

func (p *Player) GetPos() Vector {
	return p.pos
}

func (p *Player) updateAnimation() {
	p.animationFrames++
	if p.animationFrames < 2 {
		return
	}
	p.animationFrames = 0

	p.animationIdx++
	if p.animationIdx > len(p.sprites[p.state])-1 {
		p.animationIdx = 0
	}
}

func (p *Player) setState(state PlayerState) {
	if p.state == state {
		return
	}

	p.state = state
	p.animationFrames = 0
	p.animationIdx = 0
}

func (p *Player) loadSprites() {
	p.sprites[PlayerStateIdle] = []*ebiten.Image{
		utils.ScaleImage(p.loadImage(images.SantaIdle1), 0.15, 0.15),
		utils.ScaleImage(p.loadImage(images.SantaIdle2), 0.15, 0.15),
		utils.ScaleImage(p.loadImage(images.SantaIdle3), 0.15, 0.15),
		utils.ScaleImage(p.loadImage(images.SantaIdle4), 0.15, 0.15),
		utils.ScaleImage(p.loadImage(images.SantaIdle5), 0.15, 0.15),
		utils.ScaleImage(p.loadImage(images.SantaIdle6), 0.15, 0.15),
		utils.ScaleImage(p.loadImage(images.SantaIdle7), 0.15, 0.15),
		utils.ScaleImage(p.loadImage(images.SantaIdle8), 0.15, 0.15),
		utils.ScaleImage(p.loadImage(images.SantaIdle9), 0.15, 0.15),
		utils.ScaleImage(p.loadImage(images.SantaIdle10), 0.15, 0.15),
		utils.ScaleImage(p.loadImage(images.SantaIdle11), 0.15, 0.15),
		utils.ScaleImage(p.loadImage(images.SantaIdle12), 0.15, 0.15),
		utils.ScaleImage(p.loadImage(images.SantaIdle13), 0.15, 0.15),
		utils.ScaleImage(p.loadImage(images.SantaIdle14), 0.15, 0.15),
		utils.ScaleImage(p.loadImage(images.SantaIdle15), 0.15, 0.15),
		utils.ScaleImage(p.loadImage(images.SantaIdle16), 0.15, 0.15),
	}

	p.sprites[PlayerStateWalk] = []*ebiten.Image{
		utils.ScaleImage(p.loadImage(images.SantaWalk1), 0.15, 0.15),
		utils.ScaleImage(p.loadImage(images.SantaWalk2), 0.15, 0.15),
		utils.ScaleImage(p.loadImage(images.SantaWalk3), 0.15, 0.15),
		utils.ScaleImage(p.loadImage(images.SantaWalk4), 0.15, 0.15),
		utils.ScaleImage(p.loadImage(images.SantaWalk5), 0.15, 0.15),
		utils.ScaleImage(p.loadImage(images.SantaWalk6), 0.15, 0.15),
		utils.ScaleImage(p.loadImage(images.SantaWalk7), 0.15, 0.15),
		utils.ScaleImage(p.loadImage(images.SantaWalk8), 0.15, 0.15),
		utils.ScaleImage(p.loadImage(images.SantaWalk9), 0.15, 0.15),
		utils.ScaleImage(p.loadImage(images.SantaWalk10), 0.15, 0.15),
		utils.ScaleImage(p.loadImage(images.SantaWalk11), 0.15, 0.15),
		utils.ScaleImage(p.loadImage(images.SantaWalk12), 0.15, 0.15),
		utils.ScaleImage(p.loadImage(images.SantaWalk13), 0.15, 0.15),
	}
}

func (p *Player) loadImage(file []byte) *ebiten.Image {
	img, _, err := image.Decode(bytes.NewReader(file))
	if err != nil {
		log.Fatal(err)
	}

	return ebiten.NewImageFromImage(img)
}
