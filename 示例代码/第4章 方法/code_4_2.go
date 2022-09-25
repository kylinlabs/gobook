//第4章/code_4_2.go
type Point2d struct {
	Point
	y float64
}

func (p Point2d) Y() float64 {
	return p.y
}

func (p *Point2d) SetY(y float64) {
	p.y = y
}
