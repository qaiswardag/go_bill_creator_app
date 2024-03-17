package main

import (
	"bufio"
	"fmt"
	"os"
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

	name, _ := getInput("Enter your full name:", reader)

	b := newBill(name)

	fmt.Println("Created the bill ", b.name)

	return b
}

func promptOptions(b bill) {
	reader := bufio.NewReader(os.Stdin)

	opt, _ := getInput("Choose option (a: add item, s: save the bill, t: add tip): ", reader)

	fmt.Println(opt)

}

func main() {
	theBill := createBill()
	promptOptions(theBill)
}
