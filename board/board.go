package board

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	indicesOutOfBoundsError = "index out of bounds: %v"
)

type Board interface {
	IsValid(Point) bool
	Get(Point) (Tile, error)
	Set(Point, Tile) error
	Random() Point
	Width() int
	Height() int
}

type board struct {
	tiles []Tile
	n     int
	m     int
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func New(n, m int) (Board, error) {
	if n < 0 || m < 0 {
		return nil, fmt.Errorf("size is invalid: %v by %v", n, m)
	}

	tiles := make([]Tile, n*m)
	return &board{
		tiles: tiles,
		n:     n,
		m:     m,
	}, nil
}

func (this *board) IsValid(p Point) bool {
	x, y := p.Coordinates()
	return x >= 0 && x < this.n && y >= 0 && y < this.m
}

func (this *board) toInnerIndex(p Point) int {
	return p.X() + this.n*p.Y()
}

func (this *board) Get(p Point) (Tile, error) {
	if !this.IsValid(p) {
		return Invalid, fmt.Errorf(indicesOutOfBoundsError, p)
	}

	return this.tiles[this.toInnerIndex(p)], nil
}

func (this *board) Set(p Point, tile Tile) error {
	if !this.IsValid(p) {
		return fmt.Errorf(indicesOutOfBoundsError, p)
	}

	this.tiles[this.toInnerIndex(p)] = tile
	return nil
}

func (this *board) Random() Point {
	return NewPoint(rand.Intn(this.n), rand.Intn(this.m))
}

func (this *board) Width() int {
	return this.n
}

func (this *board) Height() int {
	return this.m
}
