package snake

import "snake/board"

type Snake interface {
	Len() int
	Contains(board.Point) bool
	Advance() board.Tile
	Board() board.Board
	Score() int
	AddScore(int)
	Direction() Direction
	SetDirection(Direction)
}

type snake struct {
	board     board.Board
	body      []board.Point
	head      int
	score     int
	direction Direction
}

func New(tail board.Point, directions []Direction, b board.Board) Snake {
	body := make([]board.Point, len(directions))
	body[0] = tail
	headIndex := len(directions) - 1
	for i := 0; i < headIndex; i++ {
		body[i+1] = directions[i].TranslatePoint(body[i])
	}

	for _, p := range body[:headIndex] {
		b.Set(p, board.Body)
	}

	b.Set(body[headIndex], board.Head)

	return &snake{
		board:     b,
		body:      body,
		head:      headIndex,
		direction: directions[headIndex],
	}
}

func (this *snake) Len() int {
	return len(this.body)
}

func (this *snake) Contains(p board.Point) bool {
	for _, segment := range this.body {
		if p.Equels(segment) {
			return true
		}
	}

	return false
}

func (this *snake) Board() board.Board {
	return this.board
}

func (this *snake) Score() int {
	return this.score
}

func (this *snake) AddScore(amount int) {
	this.score += amount
}

func (this *snake) tailIndex() int {
	return (this.head + 1) % this.Len()
}

func (this *snake) Advance() board.Tile {
	head := this.body[this.head]
	newHead := this.direction.TranslatePoint(head)
	b := this.Board()

	tile, err := b.Get(newHead)
	if err != nil || tile == board.Body {
		return tile
	}

	tailIndex := this.tailIndex()
	if tile == board.Apple {
		newbody := append(this.body[tailIndex:], this.body[:tailIndex]...)
		this.body = append(newbody, newHead)
		this.head = this.Len() - 1
	} else {
		tail := this.body[tailIndex]
		b.Set(tail, board.Empty)
		this.head = tailIndex
		this.body[this.head] = newHead
	}

	b.Set(head, board.Body)
	b.Set(newHead, board.Head)
	return tile
}

func (this *snake) Direction() Direction {
	return this.direction
}

func (this *snake) SetDirection(direction Direction) {
	this.direction = direction
}
