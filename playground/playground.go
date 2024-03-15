package playground

import (
	"fmt"
)

// private function
func printString (stringText string){
	fmt.Println(stringText)
}
// bits & memory
// https://pkg.go.dev/builtin

func Explore() {

	// strings
	// strings
	// strings
	// strings

	// var nameOne string = "Hello"
	// nameTwo := "world" // unable to use := outside a function
	// var nameThree string
	// fmt.Println(nameOne, nameTwo, nameThree)


	// numbers integers & bit size
	// numbers integers & bit size
	// numbers integers & bit size
	// numbers integers & bit size

	// Numeric types & bits & memory
	// https://pkg.go.dev/builtin#int8

	// var numberOne int8 = 0 // range: -128 through 127
	// var numberTwo uint8 = 0 // range: 0 through 255
	// var numberThree uint16 = 0 // range: 0 through 65535
	
	// use type `int` mostly
	// On a 64-bit system, int will typically be a 64-bit signed integer, with a range from -9223372036854775808 to 9223372036854775807.
	// var numberFour int = 922337203685477580 // the int type does not have a fixed size, its size depends on the platform

	// fmt.Println(numberOne, numberTwo, numberThree, numberFour)
	
	
	// floating number & bit size
	// floating number & bit size
	// floating number & bit size
	// floating number & bit size
	
	// For the most part, we'll be using `float64` because they offer slightly higher precision compared to `float32`. 
	// To be honest, it's not going to cause much of a memory hit on your computer.
	
	// Unlike integers we have to specify bit size for floating number which can be `float32` or `float64`
	// var numberOne float32 = 320.50
	// var numberTwo float64 = 640.50

	// default as float64
	// scoreThree :=  1.5 // default as float64
	
	// fmt.Println(numberOne, numberTwo, scoreThree)


	// Printing & Formatting Strings
	// Printing & Formatting Strings
	// Printing & Formatting Strings
	// Printing & Formatting Strings

	// fmt.Print("hello ") // does not add a new line: 
	// fmt.Print("world") // does not add a new line: 
	// result: hello word
	//
	//
	//
	// country := "Denmark"
	// city := "Copenhagen"
	// years := 20

	// // formatted strings
	// fmt.Printf("hello world I have lived in %v, %v for %v \n", country, city, years);

	//
	//
	//
	//
	//
	// Arrays
	// Arrays
	// Arrays
	// Arrays

	// var numbersOne [3]int = [3]int{10,20,30}

	// var numbersOne = [3]int{10,20,30}
	
	// In Go, arrays require elements of the same type. 
	// Therefore, you cannot directly declare an array with mixed types like int, int, string. 
	// var numbersOne = [3]interface{}{10,20,"30"}
	
	// var names = [3]string{"yoshi", "mario", "Shaun"}
	// fmt.Println(names)
	
	//
	//
	//
	//
	// Slices
	// Slices
	// Slices
	// Slices
	
	// When no number ins entered inside the [] then it create a slice
	// var data = []int{100,50,60}
	//
	// Now we can also append items to slice
	// append function do not change original Slice, but returns a new slice 
	// var dataNew = append(data, 1000)
	// fmt.Println(dataNew)


	//
	//
	//
	//
	// Standard library
	// Standard library
	// Standard library
	// Standard library
	// String: https://pkg.go.dev/strings@go1.22.1

	//  hello := "Hello world how are you doing"

	//  var containsWorld = strings.Contains(hello, "world")

	//  fmt.Println(containsWorld)


	// Uppercase
	// Uppercase
	// hello := "Hello world how are you doing"
	
	// var helloUpperCase = strings.ToUpper(hello)
	
	// fmt.Println(helloUpperCase)
	//
	//
	//
	// Split string to Array
	// Split string to Array
	// hello := "Hello world how are you doing"

	// stringToArray := strings.Split(hello, " ")

	// fmt.Println(stringToArray)

	// private function
	// private function
	// private function
	// private function
	// printString("Hello world");



	//
	//
	//
	//
	// Lops
	// Lops
	// Lops
	// Lops
}