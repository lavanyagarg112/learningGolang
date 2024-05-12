package main

import (
	"fmt"
)

type gasEngine struct { // default values set for the struct
	mpg uint8
	gallons uint8
	owner
	// owner.name can also be done


}

type electricEngine struct {
	mpkwh uint8
	kwh uint8
}

func (e gasEngine) milesLeft() uint8 { // method: like a function, except we are assigning this function to the gasengine type, so it has access to all fields
	return e.gallons*e.mpg
}

func (e electricEngine) milesLeft() uint8 { // method: like a function, except we are assigning this function to the gasengine type, so it has access to all fields
	return e.kwh*e.mpkwh
}

// to make canMakeIt more generalised - use interface

type engine interface{
	milesLeft() uint8
}

func canMakeIt(e engine, miles uint8){ // this is a normal function.. and it takes in a interface 
	// so the only requirement is that the input should have milesLeft method
	if miles <= e.milesLeft(){
		fmt.Println("you can make it there!")
	} else {
		fmt.Println("need to fuel up first!")
	}
}

// is this all like java omg
type owner struct{
	name string
}

func main(){
	fmt.Println("Structs and Interfaces")

	// structs: defining our own type - like gasEngine
	// var myEngine gasEngine = gasEngine{} // default values
	var myEngine gasEngine = gasEngine{10,20, owner{"Alex"}}
	fmt.Println(myEngine.mpg, myEngine.gallons, myEngine.owner.name)

	// anonymous struct, not reusable

	var myEngine2 = struct{
		mpg uint8
		gallons uint8
	}{25,15}

	fmt.Println(myEngine2.mpg, myEngine2.gallons)

	// structs have methods as well

	fmt.Printf("Total miles left in tank: %v", myEngine.milesLeft())
	fmt.Println()

	canMakeIt(myEngine, 50)

	var myElectric electricEngine = electricEngine{10,30}
	canMakeIt(myElectric, 100)


}