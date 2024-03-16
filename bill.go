package main

import "fmt"

type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

// make new bills
func newBill(name string) bill {
	billOjbect := bill{
		name:  name,
		items: map[string]float64{},
		tip:   0,
	}

	return billOjbect
}

// Receiver Function
// this function can only be called from bill object only
// So the function is not a function on its own
func (b *bill) formatBill() string {
	formatString := "Bill breakdown: \n"

	var total float64 = 0

	// list items
	for k, v := range b.items {
		formatString += fmt.Sprintf("%-40v $%v \n", k+":", v)

		total += v
	}

	// tip
	formatString += fmt.Sprintf("%-40v $%v\n\n", "Tip:", b.tip)

	// total
	formatString += fmt.Sprintf("%-40v $%0.2f", "Total:", total+b.tip)

	return formatString
}

// update tip
func (b *bill) updateTip(tip float64) {
	(*b).tip = tip
}

// add an item to bill
func (b *bill) addItem(name string, price float64) {
	b.items[name] = price
}
