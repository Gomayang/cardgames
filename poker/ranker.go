package poker

import (
	"sort"

	trumpcard "github.com/Gomayang/cardgames/playingcard"
)

type pairMemory map[int8]int8

func CheckRank(deck trumpcard.Deck) (*Ranking, *CardRank) {
	pairRanking, pairCardRank := pairChecker(deck)
	return pairRanking, pairCardRank
}

func flushChecker(deck trumpcard.Deck) (Rank, *trumpcard.Deck) {
	var flushDeck trumpcard.Deck
	rank := HighCard
	var flushStack int8 = 0
	sort.Slice(deck.Cards, func(i, j int) bool {
		return deck.Cards[i].Number > deck.Cards[j].Number
	})
	sort.SliceStable(deck.Cards, func(i, j int) bool {
		return deck.Cards[i].Number > deck.Cards[j].Number
	})
	flushDeck.Cards[0] = deck.Cards[0]
	for i := 1; i < len(deck.Cards); i++ {
		if flushDeck.Cards[0].Suit == deck.Cards[i].Suit {
			flushStack++
			flushDeck.Cards[flushStack] = deck.Cards[i]
			if flushStack >= 4 {
				rank = Flush
				return rank, &flushDeck
			}
		} else {
			flushStack = 0
			flushDeck.Cards = []trumpcard.Card{deck.Cards[i]}
		}
	}
	return rank, nil
}

func pairChecker(deck trumpcard.Deck) (*Ranking, *CardRank) {
	cardRank := CardRank{}
	ranking := Ranking{}
	var spades, hearts, diamonds, clovers []int8
	pairMemory := pairMemory{deck.Cards[0].Number: 0}
	for index := 1; index < len(deck.Cards); index++ {
		_, isExist := pairMemory[deck.Cards[index].Number]
		if isExist {
			pairMemory[deck.Cards[index].Number]++
		} else {
			pairMemory[deck.Cards[index].Number] = 0
		}
		switch deck.Cards[index].Suit {
		case int8(trumpcard.Spade):
			spades = append(spades, deck.Cards[index].Number)
		case int8(trumpcard.Heart):
			hearts = append(hearts, deck.Cards[index].Number)
		case int8(trumpcard.Diamond):
			diamonds = append(diamonds, deck.Cards[index].Number)
		case int8(trumpcard.Clover):
			clovers = append(clovers, deck.Cards[index].Number)
		}
	}

	//spades = [2, 4, 5, 6, 7, 10]
	var highs []int8
	var pairs []int8
	var triples []int8
	var fours []int8

	for key := range pairMemory {
		switch {
		case pairMemory[key] >= 0:
			highs = append(highs, key)
			fallthrough
		case pairMemory[key] >= 1:
			pairs = append(pairs, key)
			fallthrough
		case pairMemory[key] >= 2:
			triples = append(triples, key)
			fallthrough
		case pairMemory[key] >= 3:
			fours = append(fours, key)
		}
	}

	// Check Straight
	var straight bool = false
	var straightStack int8 = 0
	sort.Slice(highs, func(i, j int) bool {
		return highs[i] > highs[j]
	})
	var straightNumber int8 = highs[0]
	for _, number := range highs {
		if number == straightNumber-1 {
			straightStack++
			if straightStack >= 4 {
				straight = true
				break
			}
		} else {
			straightStack = 0
		}
		straightNumber = number
	}
	if straightStack == 3 && straightNumber == 2 && highs[0] == 14 {
		straight = true
	}

	// Check Ranks
	switch {
	// Check Straight Flush

	case len(fours) > 0: // Check FourCard
		sort.Slice(fours, func(i, j int) bool {
			return fours[i] > fours[j]
		})
		cardRank[FourCard] = fours[0]
		highs = elementDelete(cardRank[FourCard], highs)
		sort.Slice(highs, func(i, j int) bool {
			return highs[i] > highs[j]
		})
		cardRank[FirstCard] = highs[0]
		ranking.Rank = int8(FourCard)
		ranking.Number = cardRank[FourCard]

	case (len(triples) > 0) && (len(pairs) > 1): // Check Fullhouse
		sort.Slice(triples, func(i, j int) bool {
			return triples[i] > triples[j]
		})
		cardRank[FullHouse] = triples[0]
		sort.Slice(pairs, func(i, j int) bool {
			return pairs[i] > pairs[j]
		})
		if cardRank[FullHouse] == pairs[0] {
			cardRank[Pair] = pairs[1]
		} else {
			cardRank[Pair] = pairs[0]
		}
		ranking.Rank = int8(FullHouse)
		ranking.Number = cardRank[FullHouse]

	// Check Flushes
	case len(spades) > 4:
		sort.Slice(spades, func(i, j int) bool {
			return spades[i] > spades[j]
		})
		cardRank[Flush] = spades[0]
		cardRank[FirstCard] = spades[1]
		cardRank[SecondCard] = spades[2]
		cardRank[ThirdCard] = spades[3]
		cardRank[FourthCard] = spades[4]
		ranking.Rank = int8(Flush)
		ranking.Number = cardRank[Flush]

	case len(hearts) > 4:
		sort.Slice(hearts, func(i, j int) bool {
			return hearts[i] > hearts[j]
		})
		cardRank[Flush] = hearts[0]
		cardRank[FirstCard] = hearts[1]
		cardRank[SecondCard] = hearts[2]
		cardRank[ThirdCard] = hearts[3]
		cardRank[FourthCard] = hearts[4]
		ranking.Rank = int8(Flush)
		ranking.Number = cardRank[Flush]

	case len(diamonds) > 4:
		sort.Slice(diamonds, func(i, j int) bool {
			return diamonds[i] > diamonds[j]
		})
		cardRank[Flush] = diamonds[0]
		cardRank[FirstCard] = diamonds[1]
		cardRank[SecondCard] = diamonds[2]
		cardRank[ThirdCard] = diamonds[3]
		cardRank[FourthCard] = diamonds[4]
		ranking.Rank = int8(Flush)
		ranking.Number = cardRank[Flush]

	case len(clovers) > 4:
		sort.Slice(clovers, func(i, j int) bool {
			return clovers[i] > clovers[j]
		})
		cardRank[Flush] = clovers[0]
		cardRank[FirstCard] = clovers[1]
		cardRank[SecondCard] = clovers[2]
		cardRank[ThirdCard] = clovers[3]
		cardRank[FourthCard] = clovers[4]
		ranking.Rank = int8(Flush)
		ranking.Number = cardRank[Flush]

	// Check Straight
	case straight:
		cardRank[Straight] = straightNumber + 3
		cardRank[FirstCard] = straightNumber + 2
		cardRank[SecondCard] = straightNumber + 1
		cardRank[ThirdCard] = straightNumber
		cardRank[FourthCard] = straightNumber - 1

	// Check Triple
	case len(triples) > 0:
		sort.Slice(triples, func(i, j int) bool {
			return triples[i] > triples[j]
		})
		cardRank[Triple] = triples[0]
		sort.Slice(highs, func(i, j int) bool {
			return highs[i] > highs[j]
		})
		highs = elementDelete(cardRank[Triple], highs)
		sort.Slice(highs, func(i, j int) bool {
			return highs[i] > highs[j]
		})
		cardRank[FirstCard] = highs[0]
		cardRank[SecondCard] = highs[1]
		ranking.Rank = int8(Triple)
		ranking.Number = cardRank[Triple]

	case len(pairs) > 1: // Check TwoPair
		sort.Slice(pairs, func(i, j int) bool {
			return pairs[i] > pairs[j]
		})
		cardRank[TwoPair] = pairs[0]
		cardRank[Pair] = pairs[1]
		highs = elementDelete(cardRank[TwoPair], highs)
		highs = elementDelete(cardRank[Pair], highs)
		sort.Slice(highs, func(i, j int) bool {
			return highs[i] > highs[j]
		})
		cardRank[FirstCard] = highs[0]
		ranking.Rank = int8(TwoPair)
		ranking.Number = cardRank[TwoPair]

	case len(pairs) > 0: // Check Pair
		sort.Slice(pairs, func(i, j int) bool {
			return pairs[i] > pairs[j]
		})
		cardRank[Pair] = pairs[0]
		highs = elementDelete(cardRank[Triple], highs)
		sort.Slice(highs, func(i, j int) bool {
			return highs[i] > highs[j]
		})
		cardRank[FirstCard] = highs[0]
		cardRank[SecondCard] = highs[1]
		cardRank[ThirdCard] = highs[2]
		ranking.Rank = int8(Pair)
		ranking.Number = cardRank[Pair]

	default: // HighCard
		sort.Slice(highs, func(i, j int) bool {
			return highs[i] > highs[j]
		})
		cardRank[HighCard] = highs[0]
		cardRank[FirstCard] = highs[1]
		cardRank[SecondCard] = highs[2]
		cardRank[ThirdCard] = highs[3]
		cardRank[FourthCard] = highs[4]
		ranking.Rank = int8(HighCard)
		ranking.Number = cardRank[HighCard]
	}

	return &ranking, &cardRank
}

func elementDelete(element int8, targetSlice []int8) []int8 {
	for i := 0; i < len(targetSlice); i++ {
		if targetSlice[i] == element {
			if i == 0 {
				targetSlice = targetSlice[1:]
			} else if (0 < i) && (i < len(targetSlice)-1) {
				targetSlice = append(targetSlice[:i], targetSlice[i+1:]...)
			} else if i == len(targetSlice)-1 {
				targetSlice = targetSlice[:i]
			}
		}
	}
	return targetSlice
}
