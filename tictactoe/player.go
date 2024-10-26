package tictactoe

const (
	xSymbol = "X"
	oSymbol = "O"
	EMPTY   = " "
)

type Player interface {
	Move(b *Board)
}
