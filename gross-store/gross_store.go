package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	return map[string]int{
		"quarter_of_a_dozen": 3,
		"half_of_a_dozen":    6,
		"dozen":              12,
		"small_gross":        120,
		"gross":              144,
		"great_gross":        1728,
	}
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	return map[string]int{}
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	_, unitExists := units[unit]
	if !unitExists {
		return false
	}
	_, itemExists := bill[item]
	if itemExists {
		bill[item] += units[unit]
	} else {
		bill[item] = units[unit]
	}
	return true
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	_, unitExists := units[unit]
	if !unitExists {
		return false
	}

	_, itemExists := bill[item]
	if itemExists {
		left := bill[item] - units[unit]
		if left == 0 {
			delete(bill, item)
		} else if left < 0 {
			return false
		} else {
			bill[item] = left
		}
	} else {
		return false
	}
	return true
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	searchItem, itemExists := bill[item]
	if !itemExists {
		return 0, false
	}
	return searchItem, true
}
