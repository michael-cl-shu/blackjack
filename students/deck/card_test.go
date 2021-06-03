package deck

import (
	"testing"
)

func TestExampleCard(t *testing.T) {
	want1 := "Ace of Hearts"
	got1 := Card{Suit: Heart, Rank: Ace}.String()
	if got1 != want1 {
		t.Errorf("want: %s got: %s", want1, got1)
	}
	want2 := "Two of Spades"
	got2 := Card{Suit: Spade, Rank: Two}.String()
	if got2 != want2 {
		t.Errorf("want: %s got: %s", want2, got2)
	}
	want3 := "Joker"
	got3 := Card{Suit: Joker}.String()
	if got3 != want3 {
		t.Errorf("want: %s got: %s", want3, got3)
	}
}

func TestNew(t *testing.T) {
	cards := New()
	if len(cards) != 13*4 {
		t.Error("Wrong number of cards in a new deck")
	}
}

func TestDefaultSort(t *testing.T) {
	cards := New(DefaultSort)

	exp1st := Card{Suit: Club, Rank: Ace}
	expLast := Card{Suit: Spade, Rank: King}
	if exp1st != cards[0] {
		t.Errorf("Expected %+v Got %+v", exp1st, cards[0])
	}

	if expLast != cards[len(cards)-1] {
		t.Errorf("Expected %+v Got %+v", expLast, cards[len(cards)-1])
	}
}

func TestSort(t *testing.T) {
	cards := New(Sort(Less))

	exp1st := Card{Suit: Club, Rank: Ace}
	expLast := Card{Suit: Spade, Rank: King}
	if exp1st != cards[0] {
		t.Errorf("Expected %+v Got %+v", exp1st, cards[0])
	}

	if expLast != cards[len(cards)-1] {
		t.Errorf("Expected %+v Got %+v", expLast, cards[len(cards)-1])
	}

}

func TestShuffle(t *testing.T) {
	cards := New()
	shuffledCards := New(Shuffle)
	// fmt.Println("In Test")
	// for _, card := range shuffledCards {
	// 	fmt.Println(card)
	// }
	same := true

	for i, card := range cards {
		if card != shuffledCards[i] {
			same = false
			break
		}
	}
	if same {
		t.Error("The Deck is not shuffled")
	}
}

func TestJokers(t *testing.T) {
	cards := New(Jokers(3))
	// fmt.Printf("In Test: %d\n", len(cards))
	count := 0
	for _, c := range cards {
		if c.Suit == Joker {
			count++
		}
	}
	if count != 3 {
		t.Errorf("Expected 3. Got %d", count)
	}
}

func TestStringer(t *testing.T) {
	// Edge cases
	var r Rank = Rank(30)
	var s Suit = Suit(20)

	if r.String() != "Rank(30)" {
		t.Error("Incorrect (i Rank) String() output")
	}
	if s.String() != "Suit(20)" {
		t.Error("Incorrect (i Suit) String() output")
	}
}

func TestFilter(t *testing.T) {
	filter := func(card Card) bool {
		return card.Rank == Two || card.Rank == Three
	}
	cards := New(Filter(filter))
	for _, card := range cards {
		if card.Rank == Two || card.Rank == Three {
			t.Errorf("Rank %s and %s are not filtered", Two.String(), Three.String())
		}
	}
}

func TestDeck(t *testing.T) {
	cards := New(Deck(3))
	number := 4 * 13 * 3
	if len(cards) != number {
		t.Errorf("Expected %d Got %d", number, len(cards))
	}
}
