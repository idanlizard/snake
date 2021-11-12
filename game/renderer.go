package game

import (
	"fmt"
	"github.com/gosuri/uilive"
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
	tileChars map[board.Tile]byte
	writer    *uilive.Writer
}

func NewAsciiRenderer() Renderer {
	return &asciiRenderer{
		tileChars: map[board.Tile]byte{
			board.Empty: ' ',
			board.Body:  '#',
			board.Head:  'O',
			board.Apple: '*',
		},
		writer: uilive.New(),
	}
}

func (this *asciiRenderer) Start() {
	this.writer.Start()
}

func (this *asciiRenderer) Render(b board.Board) {
	n, m := b.Width(), b.Height()
	rowLen := n + (n-1)*spaces + 3
	rows := m + 2
	charBuffer := make([]byte, 0, rowLen*rows)

	borderBuffer := make([]byte, rowLen)
	for k := 0; k < rowLen-1; k++ {
		borderBuffer[k] = verticalBorder
	}

	borderBuffer[rowLen-1] = lineBreak
	charBuffer = append(charBuffer, borderBuffer...)

	for j := 0; j < m; j++ {
		rowBuffer := make([]byte, rowLen)

		k := 0
		rowBuffer[k] = horizontalBorder
		k++

		for i := 0; i < n-1; i++ {
			tile, _ := b.Get(board.NewPoint(i, j))
			rowBuffer[k] = this.tileChars[tile]
			k++
			for x := 0; x < spaces; x++ {
				rowBuffer[k] = space
				k++
			}
		}

		tile, _ := b.Get(board.NewPoint(n-1, j))
		rowBuffer[k] = this.tileChars[tile]
		k++
		rowBuffer[k] = horizontalBorder
		k++
		rowBuffer[k] = lineBreak
		charBuffer = append(charBuffer, rowBuffer...)
	}

	charBuffer = append(charBuffer, borderBuffer...)
	fmt.Fprint(this.writer, string(charBuffer))
	this.writer.Flush()
}

func (this *asciiRenderer) Message(s string) {
	fmt.Println(s)
}

func (this *asciiRenderer) Stop() {
	go this.writer.Stop()
}
