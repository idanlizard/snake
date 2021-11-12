package board

type Tile int

const (
	Empty Tile = iota
	Body
	Head
	Apple
	Invalid
)

var charTable = []byte{' ', '#', 'O', '*', '!'}

func (this Tile) Byte() byte {
	return charTable[int(this)]
}
