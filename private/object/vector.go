package object

type Vector struct {
	X float64
	Y float64
}

func NewVector(x, y float64) Vector {
	return Vector{
		X: x,
		Y: y,
	}
}

func (v *Vector) Add(v2 Vector) {
	v.X += v2.X
	v.Y += v2.Y
}
