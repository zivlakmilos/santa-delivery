package leveleditor

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	ScreenWidth  = 1512
	ScreenHeight = 1024
)

type LevelEditor struct {
	toolbar *Toolbar
}

func NewLevelEditor() *LevelEditor {
	return &LevelEditor{
		toolbar: NewToolbar(),
	}
}

func (l *LevelEditor) Update() error {
	return nil
}

func (l *LevelEditor) Draw(screen *ebiten.Image) {
	l.toolbar.Draw(screen)
}

func (l *LevelEditor) Layout(outsideWidth, outsideHeight int) (int, int) {
	return ScreenWidth, ScreenHeight
}

func StartLevelEditor() {
	ebiten.SetWindowSize(ScreenWidth, ScreenHeight)
	ebiten.SetWindowTitle("Santa Delivery Level Editor")
	levelEditor := NewLevelEditor()
	if err := ebiten.RunGame(levelEditor); err != nil {
		log.Fatal(err)
	}
}
