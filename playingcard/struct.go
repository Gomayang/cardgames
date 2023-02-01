package playingcard

type Suit int8

const (
	Spade Suit = iota
	Heart
	Diamond
	Clover
)

type Card struct {
	Suit   int8
	Number int8
}

type Deck struct {
	Cards []Card
}
