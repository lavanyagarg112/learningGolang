package main

// special package name - main package
// need to create function main

import (
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


}