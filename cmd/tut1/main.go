package main

// special package name - main package
// need to create function main

import (
	"errors"
	"fmt"
	"unicode/utf8"
)

// cant import / declare but not use
// so cant have any pointless vars/ imports

func main() {
	fmt.Println("Hello World!")
	// var intNum int
	// int16, int64 etc..
	var intNum2 int16 = 32767
	intNum2 = intNum2 + 1 // overflow but no compile time error, just some weird ans
	fmt.Println(intNum2)

	var floatNum float32 = 12345678.9
	// 32 precision gives back diff number
	fmt.Println(floatNum)

	// cant add a int and float etc
	// cant multiply int32 with int64
	// need to cast

	var floatNum32 float32 = 10.1
	var intNum32 int32 = 2
	var result float32 = floatNum32 + float32(intNum32)
	fmt.Println(result)

	var intNum1 int = 3
	var intNum3 int = 2
	fmt.Println(intNum1/intNum3)
	fmt.Println(intNum1%intNum3)

	var myString string = "Hello World" // single line
	// back quotes can format string on diff lines etc

	fmt.Println(myString)
	fmt.Println(len(myString)) // this is number of bytes

	// to show number of characters -
	fmt.Println(utf8.RuneCountInString(myString))
	// rune represents characters

	var myRune rune = 'a'
	fmt.Println(myRune)

	var myBoolean bool = true
	fmt.Println(myBoolean)

	// go sets default values.. eg int, string etc is 0
	// string is empty string
	// boolean is false

	myVar := "text" // shorthand - and obvious the type
	fmt.Println(myVar)

	myVar1, myVar2 := 1,2
	fmt.Println(myVar1)
	fmt.Println(myVar2)

	const myConst string = "const vale"
	// have to initialise and cannot change

	fmt.Println(myConst)

	fmt.Println("FUNCTIONS")
	printMe("HI this is me")

	var result2, remainder, err = intDivision(11, 10)
	if err != nil{
		fmt.Println(err.Error())
	} else if remainder == 0 {
		fmt.Printf("The result of the integer division is %v", result2)
	} else {
		fmt.Printf("The result is %v and the remainder is %v", result2, remainder) // printf formatting
	}
	fmt.Println()

	switch{
		case err!=nil:
			fmt.Println(err.Error())
		case remainder == 0:
			fmt.Printf("The result of the integer division is %v", result2)
		default:
			fmt.Printf("The result is %v and the remainder is %v", result2, remainder)

	}
	fmt.Println()

	// the break statement here is implied

	// can also have conditional switch

	switch remainder{
	case 0:
		fmt.Printf("The division was exact")
	case 1,2:
		fmt.Printf("The division was close")
	default:
		fmt.Printf("The division was not close ")
	}
	fmt.Println()


	fmt.Println(intDivision(4,3))


}


// functions

func printMe(printValue string){ // curly braces has to be on this line
	fmt.Println(printValue)
}

func intDivision(numerator int, denominator int) (int, int, error) { // need to specify what we want to return
	var err error // error is nil rn
	if denominator == 0 {
		err = errors.New("cannot divide by zero")
		return 0,0,err
	}
	var result int = numerator/denominator
	var remainder int = numerator % denominator
	return result, remainder, err
}