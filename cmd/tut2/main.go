package main

import "fmt"

func main(){
	var intArr [3]int32
	// fixed size of array
	// same size of array
	// indexible

	fmt.Println(intArr[0])
	fmt.Println(intArr[1:3]) // can slice also

	intArr[1] = 5

	// contiguous in memory

	fmt.Println(&intArr[0]) // to get memory location
	// + 4
	fmt.Println(&intArr[1])

	// immediate initialisation
	var intArr2 [3]int32 = [3]int32{1,2,3}
	fmt.Println(intArr2)

	intArr3 := [3]int32{1,2,3} // shorthand
	fmt.Println(intArr3)


	intArr4 := [...]int32{1,2,3} // shorthand omitting length cause fixed length alr infered
	fmt.Println(intArr4)

	var intSlice []int32 = []int32{4,5,6}
	fmt.Println(intSlice)
	fmt.Println()
	fmt.Printf("The length is %v with capacity %v", len(intSlice), cap(intSlice))
	fmt.Println()
	intSlice = append(intSlice, 7) // since overflow it increases the capacity to 6.. but length is 4
	fmt.Println(intSlice)
	fmt.Printf("The length is %v with capacity %v", len(intSlice), cap(intSlice))
	fmt.Println()

	var intSlice2 []int32 = []int32{8,9}
	intSlice = append(intSlice, intSlice2...)
	fmt.Println(intSlice)

	var intSlice3 []int32 = make([]int32, 3, 8) // third arg is the capacity... prevents program from reallocating the memory
	intSlice3[0] = 1



}