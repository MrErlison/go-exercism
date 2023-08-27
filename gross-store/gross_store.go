package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	var list map[string]int = map[string]int{}
	list["quarter_of_a_dozen"] = 3
	list["half_of_a_dozen"] = 6
	list["dozen"] = 12
	list["small_gross"] = 120
	list["gross"] = 144
	list["great_gross"] = 1728

	return list
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	return map[string]int{}
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	v, ok := units[unit]
	if ok {
		bill[item] += v
		return ok
	}

	return ok
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	quantity, ok := bill[item]
	if !ok {
		return ok
	}
	amount, ok := units[unit]
	if !ok {
		return ok
	}

	switch {
	case (quantity - amount) < 0:
		return false
	case (quantity - amount) == 0:
		delete(bill, item)
	default:
		bill[item] -= amount
	}

	return true
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	v, ok := bill[item]
	if ok {
		return v, ok
	}

	return 0, ok
}
