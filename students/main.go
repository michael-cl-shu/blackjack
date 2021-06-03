package main

import (
	"fmt"
	"strings"

	"example.com/blackjack/deck"
)

type Hand []deck.Card

func (h Hand) String() string {
	strs := make([]string, len(h))
	for i := range h {
		strs[i] = h[i].String()
	}
	return strings.Join(strs, ", ")
}

func (h Hand) DealerString() string {
	return h[0].String() + ", ***HIDDEN***"
}

func (h Hand) MinScore() int {
	score := 0
	for _, card := range h {
		score += min(int(card.Rank), 10)
	}
	return score
}

func (h Hand) Score() int {
	minScore := h.MinScore()
	if minScore > 11 {
		return minScore
	}
	// Ace in the deck, add 10 to the score
	for _, c := range h {
		if c.Rank == deck.Ace {
			return minScore + 10
		}
	}
	return minScore

}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func main() {
	// cards := deck.New(deck.Deck(3), deck.Shuffle)
	var gs GmaeState

	var card deck.Card
	var player, dealer Hand
	for i := 0; i < 2; i++ {
		for _, hand := range []*Hand{&player, &dealer} {
			card, cards = draw(cards)
			*hand = append(*hand, card)
		}
	}
	var input string
	for input != "s" {
		fmt.Println("Player:", player)
		fmt.Println("Dealer:", dealer.DealerString())
		fmt.Println("What will you do? (h)it or (s)tand?")
		fmt.Scanf("%s\n", &input)
		switch input {
		case "h":
			card, cards = draw(cards)
			player = append(player, card)
		}
	}
	// If dealer score <= 16 , dealer hits
	// If dealer has a soft 17, dealer hits
	for dealer.Score() <= 16 || (dealer.Score() == 17 && dealer.MinScore() != 17) {
		card, cards = draw(cards)
		dealer = append(dealer, card)
	}

	pScore, dScore := player.Score(), dealer.Score()
	fmt.Println("=== FINAL HANDS ===")
	fmt.Println("Player: ", player, "\nScore:", pScore)
	fmt.Println("Dealer: ", dealer, "\nScore:", dScore)

	switch {
	case pScore > 21:
		fmt.Println("You busted")
	case dScore > 21:
		fmt.Println("Dealer busted")
	case pScore > dScore:
		fmt.Println("You win!")
	case dScore > pScore:
		fmt.Println("Dealer win!")
	case pScore == dScore:
		fmt.Println("Draw")
	}
}

func draw(cards []deck.Card) (deck.Card, []deck.Card) {
	return cards[0], cards[1:]
}

type State int8

const (
	StatePlayerTurn State = iota
	StateDealerTurn
	StateHandOver
)

type GameState struct {
	Deck []deck.Card
	State
	Player Hand
	Dealer Hand
}

func clone(gs GmaeState) GameState {
	ret := GameState{
		Deck:   make([]dek.Card, len(gs.Deck)),
		State:  gs.State,
		Player: make(Hand, len(gs.Player)),
		Dealer: make(Hand, len(gs.Dealer)),
	}
	copy(ret.Deck, gs.Deck)
	copy(ret.Player, gs.Player)
	copy(ret.Dealer, gs.Player)
	return ret
}
func (gs *GameState) CurrentPlayer() *Hand {
	switch gs.State {
	case StatePlayerTurn:
		return &gs.Player
	case StateDealerTurn:
		return &gs.Dealer
	default:
		panic("it's not currently any player's turn")
	}
}
