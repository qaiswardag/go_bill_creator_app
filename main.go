package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getInput(prompt string, reader *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, error := reader.ReadString('\n')

	return strings.TrimSpace(input), error
}
func createBill() bill {
	// os.Stdin is the terminal
	reader := bufio.NewReader(os.Stdin)

	name, _ := getInput("Enter your full name: ", reader)

	b := newBill(name)

	fmt.Println("Created the bill ", b.name)

	return b
}

func promptOptions(b bill) {
	reader := bufio.NewReader(os.Stdin)

	opt, _ := getInput("Choose option (a: add item, t: add tip, s: save the bill): ", reader)

	if opt == "a" {
		name, _ := getInput("Item name: ", reader)
		price, _ := getInput("Item price: ", reader)

		priceAsNumber, err := strconv.ParseFloat(price, 64)

		if err != nil {
			fmt.Println("Price must be a number")
			promptOptions(b)
		}

		b.addItem(name, priceAsNumber)

		fmt.Println("Item added: ", name, price)

		promptOptions(b)
	}

	if opt == "t" {
		tip, _ := getInput("Enter tip amount: ", reader)

		tipAsNumber, err := strconv.ParseFloat(tip, 64)

		if err != nil {
			fmt.Println("Tip must be a number")
			promptOptions(b)
		}

		b.updateTip(tipAsNumber)

		fmt.Println("Tip added: ", tipAsNumber)

		promptOptions(b)
	}

	if opt == "s" {
		b.save()
		fmt.Println("You choose to saved the file:", b.name)
	}

	if opt == "exit" {
		fmt.Println("Exit bill")
	}

	if opt != "a" && opt != "t" && opt != "s" && opt != "exit" {
		fmt.Println("Not a valid option:", opt)
		promptOptions(b)
	}
}

func main() {
	theBill := createBill()
	promptOptions(theBill)
}
