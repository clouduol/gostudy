// Cube demonstrates embeded struct
package geometry

type Height struct{ Z float64 }

func (h *Height) Multiple(factor float64) {
	h.Z *= factor
}

func (h *Height) ScaleBy(factor float64) {
	h.Z *= factor
}

type Cube struct {
	Point
	Height
}
