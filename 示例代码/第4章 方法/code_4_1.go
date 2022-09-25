//第4章/code_4_1.go
package gom

type Point struct {
	x float64
}

func (p Point) X() float64 {
	return p.x
}

func (p *Point) SetX(x float64) {
	p.x = x
}
