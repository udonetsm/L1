package point

import (
	"math"
)

// point object
type Point struct {
	x float64
	y float64
}

// returns new point object
func NewPoint() *Point {
	return &Point{
		x: 0,
		y: 0,
	}
}

// allow to set incapsulated params for exclude methods
func (p *Point) SetParams(x, y float64) {
	p.x = x
	p.y = y
}

// count distance between two points and returns it
func Res(point1, point2 *Point) (result float64) {
	r1 := math.Pow((point1.x - point2.x), 2)
	r2 := math.Pow((point1.y - point2.y), 2)
	result = math.Sqrt(r1 + r2)
	return
}
