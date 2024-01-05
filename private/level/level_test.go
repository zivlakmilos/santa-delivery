package level_test

import (
	"bytes"
	"testing"

	"github.com/zivlakmilos/santa-delivery/private/level"
)

func TestLoadHeader(t *testing.T) {
	data := []byte{0x20, 0x24, 50, 100}
	buf := bytes.NewBuffer(data)

	level := level.NewLevel()
	level.Width = 50
	level.Height = 100
	err := level.Load(buf)
	if err != nil {
		t.Fatalf("error: %v", err)
	}

	if level.Width != 50 {
		t.Fatalf("expected buffer with len %d, but got buffer with len %d", level.Width, 50)
	}

	if level.Height != 100 {
		t.Fatalf("expected buffer with len %d, but got buffer with len %d", level.Height, 100)
	}
}

func TestSaveHeader(t *testing.T) {
	expected := []byte{0x20, 0x24, 50, 100}
	var buf bytes.Buffer

	level := level.NewLevel()
	level.Width = 50
	level.Height = 100
	err := level.Save(&buf)
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
