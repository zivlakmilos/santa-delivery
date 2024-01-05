package level_test

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/zivlakmilos/santa-delivery/private/level"
)

func TestLoadHeader(t *testing.T) {
	data := []byte{0x20, 0x24, 2, 2, 1, 4, 13, 17}
	buf := bytes.NewBuffer(data)

	l := level.NewLevel()
	err := l.Load(buf)
	if err != nil {
		t.Fatalf("error: %v", err)
	}

	if l.Width != 2 {
		t.Fatalf("expected buffer with len %d, but got buffer with len %d", l.Width, 50)
	}

	if l.Height != 2 {
		t.Fatalf("expected buffer with len %d, but got buffer with len %d", l.Height, 100)
	}
}

func TestSaveHeader(t *testing.T) {
	expected := []byte{0x20, 0x24, 2, 2, 1, 4, 13, 17}
	var buf bytes.Buffer

	l := level.NewLevel()
	l.Width = 2
	l.Height = 2
	l.Tiles = [][]*level.Tile{
		{&level.Tile{level.TileType1}, &level.Tile{level.TileType4}},
		{&level.Tile{level.TileType13}, &level.Tile{level.TileType17}},
	}
	err := l.Save(&buf)
	if err != nil {
		t.Fatalf("error: %v", err)
	}

	res := buf.Bytes()
	if len(expected) != len(res) {
		t.Fatalf("expected buffer with len %d, but got buffer with len %d", len(expected), len(res))
	}

	for idx := range expected {
		if expected[idx] != res[idx] {
			t.Fatalf("expected %d, but got %d at index %d", expected[idx], res[idx], idx)
		}
	}
}

func TestLoadTiles(t *testing.T) {
	data := []byte{0x20, 0x24, 2, 2, 1, 4, 13, 17}
	buf := bytes.NewBuffer(data)

	l := level.NewLevel()
	err := l.Load(buf)
	if err != nil {
		t.Fatalf("error: %v", err)
	}

	fmt.Printf("%v\n", l.Tiles)

	if l.Width != 2 {
		t.Fatalf("expected buffer with len %d, but got buffer with len %d", l.Width, 50)
	}

	if l.Height != 2 {
		t.Fatalf("expected buffer with len %d, but got buffer with len %d", l.Height, 100)
	}

	if l.Tiles[0][0].TileType != 1 {
		t.Fatalf("expected %d, but got %d on (%d, %d)", 1, l.Tiles[0][0], 0, 0)
	}
	if l.Tiles[0][1].TileType != 4 {
		t.Fatalf("expected %d, but got %d on (%d, %d)", 4, l.Tiles[0][1], 0, 1)
	}
	if l.Tiles[1][0].TileType != 13 {
		t.Fatalf("expected %d, but got %d on (%d, %d)", 13, l.Tiles[1][0], 1, 0)
	}
	if l.Tiles[1][1].TileType != 17 {
		t.Fatalf("expected %d, but got %d on (%d, %d)", 17, l.Tiles[1][1], 1, 1)
	}
}

func TestSaveTiles(t *testing.T) {
	expected := []byte{0x20, 0x24, 2, 2, 1, 4, 13, 17}
	var buf bytes.Buffer

	l := level.NewLevel()
	l.Width = 2
	l.Height = 2
	l.Tiles = [][]*level.Tile{
		{&level.Tile{level.TileType1}, &level.Tile{level.TileType4}},
		{&level.Tile{level.TileType13}, &level.Tile{level.TileType17}},
	}
	err := l.Save(&buf)
	if err != nil {
		t.Fatalf("error: %v", err)
	}

	res := buf.Bytes()
	if len(expected) != len(res) {
		t.Fatalf("expected buffer with len %d, but got buffer with len %d", len(expected), len(res))
	}

	for idx := range expected {
		if expected[idx] != res[idx] {
			t.Fatalf("expected %d, but got %d at index %d", expected[idx], res[idx], idx)
		}
	}
}
