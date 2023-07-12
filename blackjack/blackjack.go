package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	var price int
	switch card {
	case "ace":
		price = 11
	case "two":
		price = 2
	case "three":
		price = 3
	case "four":
		price = 4
	case "five":
		price = 5
	case "six":
		price = 6
	case "seven":
		price = 7
	case "eight":
		price = 8
	case "nine":
		price = 9
	case "ten", "jack", "queen", "king":
		price = 10
	default:
		price = 0
	}
	return price
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	var result string
	switch {
	case card1 == "ace" && card2 == "ace":
		result = "P"
	case ParseCard(card1)+ParseCard(card2) <= 11:
		result = "H"
	case (ParseCard(card1)+ParseCard(card2) <= 16 && ParseCard(card1)+ParseCard(card2) >= 12) && ParseCard(dealerCard) >= 7:
		result = "H"
	case (ParseCard(card1)+ParseCard(card2) <= 16 && ParseCard(card1)+ParseCard(card2) >= 12) && ParseCard(dealerCard) < 7:
		result = "S"
	case ParseCard(card1)+ParseCard(card2) <= 20 && ParseCard(card1)+ParseCard(card2) >= 17:
		result = "S"
	case ParseCard(card1)+ParseCard(card2) == 21 && (dealerCard == "ace" || ParseCard(dealerCard) == 10):
		result = "S"
	case ParseCard(card1)+ParseCard(card2) == 21 && (dealerCard != "ace" || ParseCard(dealerCard) != 10):
		result = "W"
	}
	return result
}
