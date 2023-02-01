package poker

type Rank int8

const (
	FirstCard Rank = iota
	SecondCard
	ThirdCard
	FourthCard
	FifthCard
	HighCard
	Pair
	TwoPair
	Triple
	Straight
	Flush
	FullHouse
	FourCard
	StraightFlush
)

type CardRank map[Rank]int8

type Ranking struct {
	Rank   int8
	Number int8
}
