package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
    return map[string]int{"quarter_of_a_dozen": 3, "half_of_a_dozen": 6, "dozen": 12, "small_gross": 120, "gross": 144, "great_gross": 1728}
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	return make(map[string]int)
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	if value, exists := units[unit]; exists == true {
        bill[item] += value

        return exists
    }

    return false
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	if gross, exists := units[unit]; exists == true {
        if qty, exists := bill[item]; exists == true {
            difference := qty - gross
            if difference > 0 {
                bill[item] = difference
                return true
            }
            if difference == 0 {
                delete(bill, item)
                return true
            }
        }
    }
    return false
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	qty, exist := bill[item]

    return qty, exist
}
