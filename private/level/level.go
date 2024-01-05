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
	Tiles [][]Tile

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

	return nil
}

func (l *Level) Save(writer io.Writer) error {
	err := l.saveHeader(writer)
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
