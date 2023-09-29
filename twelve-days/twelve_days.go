package twelve

import "fmt"

var gifts = []string{"Partridge in a Pear Tree", "Turtle Doves", "French Hens", "Calling Birds", "Gold Rings", "Geese-a-Laying", "Swans-a-Swimming", "Maids-a-Milking", "Ladies Dancing", "Lords-a-Leaping", "Pipers Piping", "Drummers Drumming"}

func Verse(i int) string {
	verse := fmt.Sprintf("On the %v day of Christmas my true love gave to me: ", getOrderedNumber(i))

	if i <= 0 || i > 12 {
		return ""
	}

	for day := i; day >= 1; day-- {
		if i == 1 {
			verse += fmt.Sprintf("%v %v.", getNumber(day), gifts[day-1])
			break
		} else if day == 1 {
			verse += fmt.Sprintf("and %v %v.", getNumber(day), gifts[day-1])
		} else {
			verse += fmt.Sprintf("%v %v, ", getNumber(day), gifts[day-1])
		}
	}

	return verse
}

func Song() string {
	days := 12
	song := ""

	for i := 1; i <= days; i++ {
		song += Verse(i)

		if i < days {
			song += "\n"
		}
	}

	return song
}

func getOrderedNumber(n int) string {
	switch n {
	case 1:
		return "first"
	case 2:
		return "second"
	case 3:
		return "third"
	case 4:
		return "fourth"
	case 5:
		return "fifth"
	case 6:
		return "sixth"
	case 7:
		return "seventh"
	case 8:
		return "eighth"
	case 9:
		return "ninth"
	case 10:
		return "tenth"
	case 11:
		return "eleventh"
	case 12:
		return "twelfth"
	}
	return ""
}
func getNumber(n int) string {
	switch n {
	case 1:
		return "a"
	case 2:
		return "two"
	case 3:
		return "three"
	case 4:
		return "four"
	case 5:
		return "five"
	case 6:
		return "six"
	case 7:
		return "seven"
	case 8:
		return "eight"
	case 9:
		return "nine"
	case 10:
		return "ten"
	case 11:
		return "eleven"
	case 12:
		return "twelve"
	}
	return ""
}
