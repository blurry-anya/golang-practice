package leap

// IsLeapYear determines if a year is leap.
func IsLeapYear(year int) bool {

	if year%4 == 0 {

		if year%400 == 0 || year%100 != 0 {
			return true
		}

	}

	return false
}
