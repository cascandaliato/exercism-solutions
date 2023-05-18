package blackjack

const (
	stand = "S"
	hit   = "H"
	split = "P"
	win   = "W"

	blackjack = 21
)

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
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
	case "ten", "jack", "queen", "king":
		return 10
	default:
		return 0
	}
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	playerValue := ParseCard(card1) + ParseCard(card2)
	dealerValue := ParseCard(dealerCard)

	switch {
	case playerValue <= 11:
		return hit
	case playerValue <= 16:
		if dealerValue < 7 {
			return stand
		}
		return hit
	case playerValue <= 20:
		return stand
	case playerValue == blackjack:
		if dealerValue >= 10 { // ace, figure or ten
			return stand
		}
		return win
	case playerValue == 22: // two aces
		return split
	default:
		return ""
	}
}
