package game

import (
	"fmt"
	term "github.com/nsf/termbox-go"
	"snake/board"
	"time"
)

const (
	spaces           = 1
	space            = ' '
	lineBreak        = '\n'
	horizontalBorder = '|'
	verticalBorder   = '-'
)

type Renderer interface {
	Start()
	Render(board.Board, *GameInfo)
	Message(string)
	Stop()
}

type asciiRenderer struct {
	tileChars map[board.Tile]rune
}

func NewAsciiRenderer() Renderer {
	return &asciiRenderer{
		tileChars: map[board.Tile]rune{
			board.Empty: ' ',
			board.Body:  '#',
			board.Head:  'O',
			board.Apple: '*',
		},
	}
}

func (this *asciiRenderer) Start() {
	term.Init()
}

func (this *asciiRenderer) Render(b board.Board, info *GameInfo) {
	n, m := b.Width(), b.Height()
	rowLen := n + (n-1)*spaces + 2

	rowNum := 0
	for k := 0; k < rowLen; k++ {
		term.SetChar(k, rowNum, verticalBorder)
	}

	rowNum++
	for j := 0; j < m; j++ {
		k := 0
		term.SetChar(k, rowNum, horizontalBorder)
		k++

		for i := 0; i < n-1; i++ {
			tile, _ := b.Get(board.NewPoint(i, j))
			term.SetChar(k, rowNum, this.tileChars[tile])
			k++
			for x := 0; x < spaces; x++ {
				term.SetChar(k, rowNum, space)
				k++
			}
		}

		tile, _ := b.Get(board.NewPoint(n-1, j))
		term.SetChar(k, rowNum, this.tileChars[tile])
		k++
		term.SetChar(k, rowNum, horizontalBorder)
		rowNum++
	}

	for k := 0; k < rowLen; k++ {
		term.SetChar(k, rowNum, verticalBorder)
	}

	rowNum += 2
	scoreStr := fmt.Sprintf("Score: %v", info.score)
	for i, char := range scoreStr {
		term.SetChar(i, rowNum, char)
	}

	rowNum++
	timeStr := fmt.Sprintf("Time:  %v", time.Time{}.Add(time.Now().Sub(info.startTime)).Format("04:05"))
	for i, char := range timeStr {
		term.SetChar(i, rowNum, char)
	}

	term.Flush()
}

func (this *asciiRenderer) Message(s string) {
	fmt.Println(s)
}

func (this *asciiRenderer) Stop() {
	term.Close()
}
