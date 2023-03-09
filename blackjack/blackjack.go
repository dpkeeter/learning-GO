package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	switch card {

	case "ace":
		return 11
	case "ten", "queen", "jack", "king":
		return 10
	case "nine":
		return 9
	case "eight":
		return 8
	case "seven":
		return 7
	case "six":
		return 6
	case "five":
		return 5
	case "four":
		return 4
	case "three":
		return 3
	case "two":
		return 2
	default:
		return 0
	}

}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	val1 := ParseCard(card1)
	val2 := ParseCard(card2)
	val3 := ParseCard(dealerCard)
	if card1 == "ace" && card2 == "ace" {
		return "P"
	}
	if val1+val2 == 21 {
		if val3 > 9 {
			return "S"
		}
		return "W"
	}
	if val1+val2 >= 17 {
		return "S"
	}
	if val1+val2 >= 12 {
		if val3 >= 7 {
			return "H"
		}
		return "S"
	}
	return "H"
}
