package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	/*
		// Better way, but not this lesson
		var cards = map[string]int{
			"ace": 11, "two": 2, "three": 3, "four": 4, "five": 5, "six": 6, "seven": 7,
			"eight": 8, "nine": 9, "ten": 10, "jack": 10, "queen": 10, "king": 10,
		}

		value, ok := cards[card]
		if ok {
			return value
		}
		return 0
	*/
	switch card {
	case "ace":
		return 11
	case "two":
		return 2
	case "three":
		return 3
	case "four":
		return 4
	case "five":
		return 5
	case "six":
		return 6
	case "seven":
		return 7
	case "eight":
		return 8
	case "nine":
		return 9
	case "ten":
		return 10
	case "jack":
		return 10
	case "queen":
		return 10
	case "king":
		return 10
	default:
		return 0
	}
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	var strategy string
	sumUp := ParseCard(card1) + ParseCard(card2)

	switch {
	case sumUp <= 11:
		strategy = "H"
	case 12 <= sumUp && sumUp <= 16:
		if ParseCard(dealerCard) >= 7 {
			strategy = "H"
		} else {
			strategy = "S"
		}
	case 17 <= sumUp && sumUp <= 20:
		strategy = "S"
	case sumUp == 21:
		if ParseCard(dealerCard) < 10 {
			strategy = "W"
		} else {
			strategy = "S"
		}
	case sumUp == 22:
		strategy = "P"
	}

	return strategy
}
