package level

type TileType byte

const (
	TileTypeNone TileType = 0
	TileType1    TileType = 1
	TileType2    TileType = 2
	TileType3    TileType = 3
	TileType4    TileType = 4
	TileType5    TileType = 5
	TileType6    TileType = 6
	TileType7    TileType = 7
	TileType8    TileType = 8
	TileType9    TileType = 9
	TileType10   TileType = 10
	TileType11   TileType = 11
	TileType12   TileType = 12
	TileType13   TileType = 13
	TileType14   TileType = 14
	TileType15   TileType = 15
	TileType16   TileType = 16
	TileType17   TileType = 17
	TileType18   TileType = 18
)

type Tile struct {
	TileType TileType
}

func NewTile(tileType TileType) *Tile {
	return &Tile{
		TileType: tileType,
	}
}
