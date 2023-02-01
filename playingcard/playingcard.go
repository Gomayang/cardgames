package playingcard

import (
	"math/rand"
	"time"
)

// NewDeck create a new standard deck.
func NewDeck() *Deck {
	deck := Deck{}
	return &deck
}

// AddDeck return new ordered trump decks.
// You can choose amount of deck.
func (deck *Deck) AddDeck(amount int) {
	for i := 0; i < amount; i++ {
		var suit int8
		for suit = 0; suit < 4; suit++ {
			var number int8
			for number = 2; number < 15; number++ {
				deck.Cards = append(deck.Cards, Card{suit, number})
			}
		}
	}
}

// Shuffle the deck
func (deck *Deck) Shuffle() {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(deck.Cards), func(i, j int) {
		deck.Cards[i], deck.Cards[j] = deck.Cards[j], deck.Cards[i]
	})
}

// Clear the deck
func (deck *Deck) Clear() {
	deck.Cards = nil
}

// Remove Card deck.Cards[index]
func (deck *Deck) Remove(index int) {
	if index == 0 {
		deck.Cards = deck.Cards[1:]
	} else if (0 < index) && (index < len(deck.Cards)-1) {
		deck.Cards = append(deck.Cards[:index], deck.Cards[index+1:]...)
	} else if index == len(deck.Cards)-1 {
		deck.Cards = deck.Cards[:index]
	}

}

// Draw Cards.
// It returns drawed deck
func (deck *Deck) Draw(count int) *Deck {
	if count > len(deck.Cards) {
		return nil
	}
	drawed := Deck{}
	drawed.Cards = deck.Cards[:count]
	deck.Cards = deck.Cards[count:]
	return &drawed
}
