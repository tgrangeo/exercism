package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	store := make(map[string]int,6)
    store["quarter_of_a_dozen"] = 3
    store["half_of_a_dozen"] = 6
	store["dozen"] = 12 
	store["small_gross"] = 120
	store["gross"] = 144
	store["great_gross"] = 1728
	return store
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	bill := make(map[string]int)
    return bill
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
    _, exists := units[unit]
    if exists == false{
        return false
    }
	_, exists = bill[item]
    if exists == true {
        bill[item] += units[unit]
    } else {
    	bill[item] = units[unit]
    }
	return true
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	_, exists := units[unit]
    if exists == false {
        return false
    }
    _, exists = bill[item]
    if exists == true {
        if (bill[item] - units[unit] < 0){
            return false
        }
        bill[item] -= units[unit]
        if (bill[item] == 0){
        	delete(bill,item)
        }
    } else {
    	return false
    }
	return true
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	_ , exists := bill[item]
    if exists == true {
        return bill[item],true
    } else{
		return 0 , false	
    }
}
