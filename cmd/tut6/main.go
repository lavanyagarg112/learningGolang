package main

import "fmt"

func main() {
	fmt.Println("Pointers")

	var p *int32 = new(int32) // this will store pointer/ memory address which will take up 32/64 bits depending on system
	// this pointer is nil, it doesnt have an address assigned to it yet
	// it stores the memory location only when we do new(int32)

	// if we dont do that, we get null pointer exception on runtime
	// so assigning gives null pointer so the *p all give null pointers
	// p itself prints out nil

	var i int32 // this is int storing default 0.. this stores value at the address

	fmt.Printf("P itself is: %v", p)
	fmt.Printf("The value p points to is: %v", *p) // pointer stores default values itself which is accessed using star
	fmt.Printf("\nThe value of i is: %v", i)

	*p = 10 // now p points to another memory location which stores 10

	p = &i // p now refers to the memory address of i, now both p, and i refer to the same
	*p = 1 // now i's value is also changes

	fmt.Printf("\n The value p points to is: %v", *p)
	fmt.Printf("\n The value of i is: %v", i)

	// different from using a normal variable
	// changing k's value again doesnt change i's value
	var k int32 = 2
	i = k
	k = 3
	fmt.Println()
	fmt.Println(i)
	fmt.Println(k)

	// the only exception in the non pointer behaviour for copying is in slicing

	// this is because is slicing we are simply copying the pointers
	// omg this is like the python thing

	var slice = []int32{1,2,3}
	var sliceCopy = slice
	sliceCopy[2] = 4
	fmt.Println(slice)
	fmt.Println(sliceCopy)

	fmt.Println("Seeing how pointers work in functions")

	var thing1 = [5]float64{1,2,3,4,5}
	fmt.Printf("\n The memort location of thing1 array is: %p", &thing1)
	var result [5]float64 = square(thing1)
	fmt.Printf("\nThe result is: %v", result)


	var thing11 = [5]float64{1,2,3,4,5}
	fmt.Printf("\n The memort location of thing1 array is: %p", &thing11)
	square2(&thing11)
	fmt.Printf("\nThe result is: %v", thing11) // since passed pointers, the array was modified already!

}


func square(thing2 [5]float64) [5]float64 {
	fmt.Printf("\n The memort location of thing2 array is: %p", &thing2)
	// these are two diff arrays
	// creating a copy
	// using way more memory needed
	// instead user pointers!

	for i := range thing2{
		thing2[i] = thing2[i]*thing2[i]
	}
	return thing2
}

// taking in a pointer instead

func square2(thing2 *[5]float64) [5]float64 {
	fmt.Printf("\n The memory location of thing2 array is: %p", thing2)

	for i := range thing2{
		thing2[i] = thing2[i]*thing2[i]
	}
	return *thing2
}