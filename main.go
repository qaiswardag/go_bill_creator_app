package main

import "fmt"

func main() {

	theBill := newBill("mario's bill")

	fmt.Println(theBill.formatBill())
}
