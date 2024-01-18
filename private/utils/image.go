package utils

import "github.com/hajimehoshi/ebiten/v2"

func FlipImageHorizontal(source *ebiten.Image) *ebiten.Image {
	result := ebiten.NewImage(source.Bounds().Dx(), source.Bounds().Dy())
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(-1, 1)
	op.GeoM.Translate(float64(source.Bounds().Dx()), 0)
	result.DrawImage(source, op)
	return result
}

func FlipImageVertical(source *ebiten.Image) *ebiten.Image {
	result := ebiten.NewImage(source.Bounds().Dx(), source.Bounds().Dy())
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(0, -1)
	op.GeoM.Translate(0, float64(source.Bounds().Dy()))
	result.DrawImage(source, op)
	return result
}
