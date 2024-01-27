package leveleditor

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/zivlakmilos/santa-delivery/private/level"
	"github.com/zivlakmilos/santa-delivery/private/utils"
	"github.com/zivlakmilos/santa-delivery/resources/images"
)

type Tool struct {
	img      *ebiten.Image
	tileType level.TileType
	x        float64
	y        float64
}

type Toolbar struct {
	tools []*Tool
}

func NewToolbar() *Toolbar {
	t := &Toolbar{}
	t.loadTiles()
	return t
}

func (t *Toolbar) Update() error {
	return nil
}

func (t *Toolbar) Draw(screen *ebiten.Image) {
	for _, tool := range t.tools {
		opt := &ebiten.DrawImageOptions{}
		opt.GeoM.Translate(tool.x, tool.y)
		screen.DrawImage(tool.img, opt)

		x, y := ebiten.CursorPosition()
		if x >= int(tool.x) && x <= int(tool.x)+tool.img.Bounds().Dx() &&
			y >= int(tool.y) && y <= int(tool.y)+tool.img.Bounds().Dy() {
			vector.StrokeRect(screen, float32(tool.x), float32(tool.y), 63, 63, 2, color.RGBA{R: 0, G: 255, B: 0}, false)
		}
	}
}

func (t *Toolbar) loadTiles() {
	t.addTool(0, level.TileTypeNone, utils.ScaleImage(utils.CreateBorderImage(128, 128, 10, color.White), 0.5, 0.5))
	t.addTool(1, level.TileType1, utils.ScaleImage(utils.LoadImage(images.Tile1Png), 0.5, 0.5))
	t.addTool(2, level.TileType2, utils.ScaleImage(utils.LoadImage(images.Tile2Png), 0.5, 0.5))
	t.addTool(3, level.TileType3, utils.ScaleImage(utils.LoadImage(images.Tile3Png), 0.5, 0.5))
	t.addTool(4, level.TileType4, utils.ScaleImage(utils.LoadImage(images.Tile4Png), 0.5, 0.5))
	t.addTool(5, level.TileType5, utils.ScaleImage(utils.LoadImage(images.Tile5Png), 0.5, 0.5))
	t.addTool(6, level.TileType6, utils.ScaleImage(utils.LoadImage(images.Tile6Png), 0.5, 0.5))
	t.addTool(7, level.TileType7, utils.ScaleImage(utils.LoadImage(images.Tile7Png), 0.5, 0.5))
	t.addTool(8, level.TileType8, utils.ScaleImage(utils.LoadImage(images.Tile8Png), 0.5, 0.5))
	t.addTool(9, level.TileType9, utils.ScaleImage(utils.LoadImage(images.Tile9Png), 0.5, 0.5))
	t.addTool(10, level.TileType10, utils.ScaleImage(utils.LoadImage(images.Tile10Png), 0.5, 0.5))
	t.addTool(11, level.TileType11, utils.ScaleImage(utils.LoadImage(images.Tile11Png), 0.5, 0.5))
	t.addTool(12, level.TileType12, utils.ScaleImage(utils.LoadImage(images.Tile12Png), 0.5, 0.5))
	t.addTool(13, level.TileType13, utils.ScaleImage(utils.LoadImage(images.Tile13Png), 0.5, 0.5))
	t.addTool(14, level.TileType14, utils.ScaleImage(utils.LoadImage(images.Tile14Png), 0.5, 0.5))
	t.addTool(15, level.TileType15, utils.ScaleImage(utils.LoadImage(images.Tile15Png), 0.5, 0.5))
	t.addTool(16, level.TileType16, utils.ScaleImage(utils.LoadImage(images.Tile16Png), 0.5, 0.5))
	t.addTool(17, level.TileType17, utils.ScaleImage(utils.LoadImage(images.Tile17Png), 0.5, 0.5))
	t.addTool(18, level.TileType18, utils.ScaleImage(utils.LoadImage(images.Tile18Png), 0.5, 0.5))
}

func (t *Toolbar) addTool(idx int, tileYype level.TileType, img *ebiten.Image) {
	tool := &Tool{
		img: img,
		x:   10 + float64(idx)*75,
		y:   10,
	}

	t.tools = append(t.tools, tool)
}
