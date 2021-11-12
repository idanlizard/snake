package snake

import (
	"snake/board"
)

type Direction string

const (
	Up    Direction = "up"
	Down  Direction = "down"
	Left  Direction = "left"
	Right Direction = "right"
)

func (this Direction) String() string {
	return string(this)
}

func (this Direction) TranslatePoint(p board.Point) board.Point {
	x, y := p.Coordinates()
	switch this {
	case Up:
		y--
	case Down:
		y++
	case Left:
		x--
	case Right:
		x++
	}

	return board.NewPoint(x, y)

}

func (this Direction) IsParallelTo(d Direction) bool {
	return (this == Up || this == Down) == (d == Up || d == Down)
}
