package object

type Camera struct {
	pos    Vector
	width  float64
	height float64
}

func NewCamera(x, y, width, height float64) *Camera {
	return &Camera{
		pos:    NewVector(x, y),
		width:  width,
		height: height,
	}
}

func (c *Camera) MoveTo(pos Vector) {
	c.pos = pos

	if c.pos.X < 0 {
		c.pos.X = 0
	}
	if c.pos.Y < 0 {
		c.pos.Y = 0
	}
}

func (c *Camera) GetPos() Vector {
	return c.pos
}
