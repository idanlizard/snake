package game

import (
	term "github.com/nsf/termbox-go"
	"snake/board"
	"snake/snake"
	"sync"
	"time"
)

const appleScore = 10

var (
	keyToDirection = map[term.Key]snake.Direction{
		term.KeyArrowUp:    snake.Up,
		term.KeyArrowDown:  snake.Down,
		term.KeyArrowLeft:  snake.Left,
		term.KeyArrowRight: snake.Right,
	}
)

type Game interface {
	Start()
	End()
}

type game struct {
	snek      snake.Snake
	board     board.Board
	stop      chan struct{}
	frameTime <-chan struct{}
	once      sync.Once
	renderer  Renderer
	info      *GameInfo
	fps       int
}

func (this *game) end() {
	this.renderer.Stop()
	this.stop <- struct{}{}
	close(this.stop)
}

func (this *game) End() {
	this.once.Do(this.end)
}

func New(n, m, fps int) (Game, error) {
	b, err := board.New(n, m)
	if err != nil {
		return nil, err
	}

	s := snake.New(board.NewPoint(0, m/2), []snake.Direction{snake.Right, snake.Right, snake.Right}, b)

	return &game{
		board:    b,
		snek:     s,
		stop:     make(chan struct{}, 1),
		renderer: NewAsciiRenderer(),
		info:     new(GameInfo),
		fps:      fps,
	}, nil
}

func (this *game) addApple() {
	for {
		apple := this.board.Random()
		if this.snek.Contains(apple) {
			continue
		}

		this.board.Set(apple, board.Apple)
		return
	}

}

func (this *game) frame() {
	switch this.snek.Advance() {
	case board.Apple:
		this.info.score += appleScore
		this.addApple()

	case board.Invalid, board.Body:
		this.End()
		return
	}

	this.renderer.Render(this.board, this.info)
}

func (this *game) Start() {
	this.addApple()
	this.info.startTime = time.Now()
	this.renderer.Render(this.board, this.info)

	timer := time.NewTicker(time.Second / time.Duration(this.fps))

	listener := NewKeyboardListener()
	defer listener.Stop()

	newDirection := this.snek.Direction()
	for {
		select {
		case <-this.stop:
			timer.Stop()
			return

		case <-timer.C:
			this.snek.SetDirection(newDirection)
			this.frame()

		case key := <-listener.Listen():
			direction, ok := keyToDirection[key]
			if !ok {
				break
			}

			if direction.IsParallelTo(this.snek.Direction()) {
				break
			}

			newDirection = direction
		}
	}
}
