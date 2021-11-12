package board

import "fmt"

type Point interface {
	X() int
	Y() int
	Coordinates() (int, int)
	Equels(Point) bool
}

type point struct {
	x int
	y int
}

func NewPoint(x, y int) Point {
	return &point{x, y}
}

func (this *point) X() int {
	return this.x
}

func (this *point) Y() int {
	return this.y
}

func (this *point) Coordinates() (int, int) {
	return this.X(), this.Y()
}

func (this *point) String() string {
	return fmt.Sprintf("(%v, %v)", this.x, this.y)
}

func (this *point) Equels(p Point) bool {
	return this.x == p.X() && this.y == p.Y()
}
