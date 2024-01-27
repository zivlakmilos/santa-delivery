package utils

import (
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
	"golang.org/x/image/font/gofont/goregular"
)

func CreateFont(ttf []byte, size float64) (font.Face, error) {
	font, err := truetype.Parse(ttf)
	if err != nil {
		return nil, err
	}

	return truetype.NewFace(font, &truetype.Options{Size: size}), nil
}

func CreateDefaultFont(size float64) (font.Face, error) {
	return CreateFont(goregular.TTF, size)
}
