package shape

type Rect struct {
	l, b float32
}

func (r *Rect) Area() float32 {
	return r.l * r.b
}
func Areaof(r *Rect) float32 {
	return r.l * r.b
}
