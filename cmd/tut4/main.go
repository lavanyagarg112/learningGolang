package main

import (
	"fmt"
	"strings"
)

func main(){

	fmt.Println("Strings, runes and bytes")

	var myString = "résumé"
	var indexed = myString[0]
	fmt.Println(myString)
	fmt.Printf("%v, %T\n", indexed, indexed) // we get a number and unsigned int8 i.e. uint8 on printing value, type

	for i,v := range myString {
		fmt.Println(i,v)
	}

	// This prints -
	// 114, uint8
	// 0 114
	// 1 233
	// 3 115
	// 4 117
	// 5 109
	// 6 233
	// why???

	// go represents strings as utf8 coding
	// string as binary 

	// before ascii - 7 bits
	// but what about emojis etc?
	// extend bits -> utf 32
	// utf 8 uses variable length encoding and uses pre defined encoding pattern
	// we can see how many bytes it uses by the prefix
	// so standard ascii is 0
	// then based on 110 etc

	// so now 114 is value of r
	// then é has 2 bytes so skips index 2

	// that is why now length of myString is 8
	// string is array of bytes

	fmt.Println(len(myString))

	// Using runes now

	var myString2 = []rune("résumé") // casting to array of runes
	// runes are just alias for int32
	var indexed2 = myString2[1]
	fmt.Println(myString2)
	fmt.Printf("%v, %T\n", indexed2, indexed2) // we 233 and int32 now

	for i,v := range myString2 {
		fmt.Println(i,v)
	}
	// get contiguous index

	fmt.Println(len(myString2))

	// can use runes using single quote as well
	var myRune = 'a'
	fmt.Printf("\nmyRune = %v", myRune)

	// String building

	var strSlice = []string{"s", "u", "b", "w", "a", "y"}
	var catStr = ""
	for i := range strSlice{ // takes index only
		catStr += strSlice[i]
	}
	fmt.Printf("\n%v", catStr)

	// string are immutable in go
	// so cannot do strSlice[1] = something else

	// concatenate is creating a new string everytime
	// so better to use string builder

	var strBuilder strings.Builder
	for i := range strSlice{
		strBuilder.WriteString(strSlice[i])
	}

	var catStr2 = strBuilder.String()
	fmt.Printf("\n%v", catStr2)




}