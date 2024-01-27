package ui

import (
	"fmt"
	"image/color"
	"log"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/zivlakmilos/santa-delivery/private/utils"
	"golang.org/x/image/font"
)

type Text struct {
	value    string
	focused  bool
	x        float32
	y        float32
	width    float32
	height   float32
	fontSize float32
	font     font.Face
}

func NewText(x, y, width float32, fontSize float64, value string) *Text {
	t := &Text{
		x:        x,
		y:        y,
		width:    width,
		height:   float32(fontSize) + 10,
		fontSize: float32(fontSize),
		value:    value,
	}
	t.loadFonts(fontSize)
	return t
}

func (t *Text) Update() error {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		if float32(x) >= t.x && float32(x) <= t.x+t.width && float32(y) >= t.y && float32(y) <= t.y+t.height {
			t.focused = true
		} else {
			t.focused = false
		}
	}

	if t.focused {
		keys := []ebiten.Key{}
		keys = inpututil.AppendJustPressedKeys(keys)
		for _, key := range keys {
			strKey := key.String()
			switch strKey {
			case "Backspace":
				if len(t.value) > 0 {
					t.value = t.value[:len(t.value)-1]
				}
				continue
			case "Space":
				t.value += " "
				continue
			}

			if strings.HasPrefix(strKey, "Numpad") || strings.HasPrefix(strKey, "Digit") {
				strKey = string(strKey[len(strKey)-1])
				fmt.Printf("%v\n", strKey)
			}

			if len(strKey) != 1 {
				continue
			}
			keyCode := strKey[0]
			if (keyCode >= 'A' && keyCode <= 'Z') || (keyCode >= 'a' && keyCode <= 'z') || (keyCode >= '0' && keyCode <= '9') {
				t.value += strKey
			}
		}
	}

	return nil
}

func (t *Text) Draw(screen *ebiten.Image) {
	var border color.Color = color.White
	if t.focused {
		border = color.RGBA{R: 0, G: 255, B: 0}
	}

	text.Draw(screen, t.value, t.font, int(t.x)+10, int(t.y+t.fontSize), color.White)
	vector.StrokeRect(screen, t.x, t.y, t.width, t.height, 2, border, false)
}

func (t *Text) SetFocus(focused bool) {
	t.focused = focused
}

func (t *Text) SetValue(value string) {
	t.value = value
}

func (t *Text) GetValue() string {
	return t.value
}

func (t *Text) GetPos() (float32, float32) {
	return t.x, t.y
}

func (t *Text) loadFonts(fontSize float64) {
	f, err := utils.CreateDefaultFont(fontSize)
	if err != nil {
		log.Fatal(err)
	}

	t.font = f
}
