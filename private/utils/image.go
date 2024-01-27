package utils

import (
	"bytes"
	"image"
	"image/color"
	"log"

	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

func LoadImage(file []byte) *ebiten.Image {
	img, _, err := image.Decode(bytes.NewReader(file))
	if err != nil {
		log.Fatal(err)
	}

	return ebiten.NewImageFromImage(img)
}

func CreateFillImage(width, height int, color color.Color) *ebiten.Image {
	result := ebiten.NewImage(width, height)
	vector.DrawFilledRect(result, 0, 0, float32(result.Bounds().Dx()), float32(result.Bounds().Dy()), color, false)
	return result
}

func CreateBorderImage(width, height int, w float32, color color.Color) *ebiten.Image {
	result := ebiten.NewImage(width, height)
	vector.StrokeRect(result, 0, 0, float32(result.Bounds().Dx()), float32(result.Bounds().Dy()), w, color, false)
	return result
}

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

func ScaleImage(source *ebiten.Image, x, y float64) *ebiten.Image {
	result := ebiten.NewImage(int(float64(source.Bounds().Dx())*x), int(float64(source.Bounds().Dy())*y))
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(x, y)
	result.DrawImage(source, op)
	return result
}
