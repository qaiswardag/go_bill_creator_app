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
		name: name,
		items: map[string]float64{
			"pie": 5.99, "cake": 3.99,
		},
		tip: 0,
	}

	return billOjbect
}

// Receiver Function
// this function can only be called from bill object only
// So the function is not a function on its own
func (b bill) formatBill() string {
	formatString := "Bill breakdown: \n"

	var total float64 = 0

	// list items
	for k, v := range b.items {
		formatString += fmt.Sprintf("%-25v $%v \n", k+":", v)

		total += v
	}

	// total
	formatString += fmt.Sprintf("%-25v $%0.2f", "total:", total)

	return formatString
}
