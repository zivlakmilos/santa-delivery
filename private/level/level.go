package level

import (
	"fmt"
	"io"
)

const (
	magic1 = 0x20
	magic2 = 0x24
)

type Level struct {
	Tiles [][]*Tile

	Width  int
	Height int
}

func NewLevel() *Level {
	return &Level{}
}

func (l *Level) Load(reader io.Reader) error {
	err := l.loadHeader(reader)
	if err != nil {
		return err
	}

	err = l.loadTiles(reader)
	if err != nil {
		return err
	}

	return nil
}

func (l *Level) Save(writer io.Writer) error {
	err := l.saveHeader(writer)
	if err != nil {
		return err
	}

	err = l.saveTiles(writer)
	if err != nil {
		return err
	}

	return nil
}

func (l *Level) loadHeader(reader io.Reader) error {
	buf := make([]byte, 4)

	n, err := reader.Read(buf)
	if err != nil {
		return err
	}
	if n < 2 {
		return fmt.Errorf("wrong file type")
	}
	if buf[0] != magic1 || buf[1] != magic2 {
		return fmt.Errorf("wrong file type")
	}

	if n < 4 {
		return fmt.Errorf("missing width and height")
	}

	l.Width = int(buf[2])
	l.Height = int(buf[3])

	return nil
}

func (l *Level) loadTiles(reader io.Reader) error {
	bufLen := l.Width * l.Height
	buf := make([]byte, bufLen)

	n, err := reader.Read(buf)
	if err != nil {
		return err
	}

	if n != bufLen {
		return fmt.Errorf("error while loading tiles")
	}

	l.Tiles = make([][]*Tile, l.Height)
	for r := 0; r < l.Height; r++ {
		l.Tiles[r] = make([]*Tile, l.Width)
		for c := 0; c < l.Width; c++ {
			tile := NewTile(TileType(buf[r*l.Width+c]))
			l.Tiles[r][c] = tile
		}
	}

	return nil
}

func (l *Level) saveHeader(writer io.Writer) error {
	buf := []byte{0x20, 0x24, byte(l.Width), byte(l.Height)}
	n, err := writer.Write(buf)
	if err != nil {
		return err
	}

	if n != 4 {
		return fmt.Errorf("error while saving header")
	}

	return nil
}

func (l *Level) saveTiles(writer io.Writer) error {
	for r := 0; r < l.Height; r++ {
		row := make([]byte, l.Width)
		for c := 0; c < l.Width; c++ {
			tile := l.Tiles[r][c]
			row[c] = byte(tile.TileType)
		}
		_, err := writer.Write(row)
		if err != nil {
			return err
		}
	}

	return nil
}
