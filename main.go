package main

import "fmt"

func main() {

	theBill := newBill("mario's bill")
	theBill.updateTip(10)

	theBill.addItem("Fried Egg", 8)
	theBill.addItem("Steak with Avocado", 20)
	theBill.addItem("Ice cream", 10)

	fmt.Println(theBill.formatBill())
}
