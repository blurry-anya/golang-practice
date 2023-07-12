package twofer

import "fmt"

func ShareWith(name string) string {
	if len(name) > 1 {
		return fmt.Sprintf("One for %v, one for me.", name)
	}
	return "One for you, one for me."
}
