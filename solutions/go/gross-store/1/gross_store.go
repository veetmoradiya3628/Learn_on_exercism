package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
    units := map[string]int{}; 
    units["quarter_of_a_dozen"] = 3
    units["half_of_a_dozen"] = 6
    units["dozen"] = 12
    units["small_gross"] = 120
    units["gross"] = 144
    units["great_gross"] = 1728
    return units;
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	bills := map[string]int{}
    return bills;
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	unitValue, exists := units[unit]
    if !exists {
        return false
    }
	bill[item] += unitValue;
    return true;
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	item_quantity, is_item_exists := bill[item]
	if !is_item_exists {
		return false
	}
	unit_size, is_unit_exists := units[unit]
	if !is_unit_exists {
		return false
	}
	if item_quantity < unit_size {
		return false
	}
	if item_quantity == unit_size {
		delete(bill, item)
	} else {
		bill[item] -= unit_size
	}
	return true
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	item_quantity, is_item_exists := bill[item]
    if !is_item_exists {
        return 0, false
    }
    return item_quantity, true
}