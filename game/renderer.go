package game

import (
	"fmt"
	term "github.com/nsf/termbox-go"
	"snake/board"
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
	Render(board.Board)
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

func (this *asciiRenderer) Render(b board.Board) {
	n, m := b.Width(), b.Height()
	rowLen := n + (n-1)*spaces + 3

	rowNum := 0
	for k := 0; k < rowLen-1; k++ {
		term.SetChar(k, rowNum, verticalBorder)
	}

	term.SetChar(rowLen-1, rowNum, lineBreak)
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
		k++
		term.SetChar(k, rowNum, lineBreak)
		rowNum++
	}

	for k := 0; k < rowLen-1; k++ {
		term.SetChar(k, rowNum, verticalBorder)
	}

	term.SetChar(rowLen-1, rowNum, lineBreak)

	rowNum += 2
	term.Flush()
}

func (this *asciiRenderer) Message(s string) {
	fmt.Println(s)
}

func (this *asciiRenderer) Stop() {
	term.Close()
}
