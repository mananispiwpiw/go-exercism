package gross

//import "fmt"

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	var units = map[string]int{
		"quarter_of_a_dozen": 3,
		"half_of_a_dozen":    6,
		"dozen":              12,
		"small_gross":        120,
		"gross":              144,
		"great_gross":        1728,
	}

	return units
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	var newBill = make(map[string]int)
	return newBill
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	value, exists := units[unit]
	if exists {
		bill[item] += value
		return true
	}
	return false
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	//Check item in the bill or not
	value, exists := units[unit]
	if !exists {
		return false
	}

	//Make sure that quantitity is more than 0
	if bill[item] >= value {
		//If condition occured, reduce the quantity
		bill[item] -= value
		//And if the bill map is equal to zero, it will delete the map
		if bill[item] == 0 {
			delete(bill, item)
		}
		return true
	}

	//If none of them condition passed, it means invalid.
	return false
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	value, exists := bill[item]
	if exists {
		return value, true
	}
	return 0, false
}
