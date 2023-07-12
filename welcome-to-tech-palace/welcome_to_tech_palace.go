package techpalace

import (
	"fmt"
	"strings"
)

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	return fmt.Sprintf("Welcome to the Tech Palace, %v", strings.ToUpper(customer))
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	border := strings.Repeat("*", numStarsPerLine)
	return fmt.Sprintf("%v\n%v\n%v", border, welcomeMsg, border)
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	oldMsg = strings.TrimFunc(oldMsg, func(r rune) bool {
		return r == '*'
	})
	oldMsg = strings.ReplaceAll(oldMsg, "*", "")
	oldMsg = strings.TrimSpace(oldMsg)
	return oldMsg
}
