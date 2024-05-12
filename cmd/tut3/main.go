package main

import "fmt"

func main(){

	fmt.Println("Map")

	// key value pairs

	var myMap map[string]uint8 = make(map[string]uint8)
	fmt.Println(myMap)

	var myMap2 = map[string]uint8{"Adam":23, "Sarah":45}
	fmt.Println(myMap2["Adam"])
	fmt.Println(myMap2["Jason"]) // even though value doesnt exist, go prints a default value

	var age, ok = myMap2["Jason"] // second argument is boolean based on if key exists or not

	// delete (myMap2, "Adam")

	if ok{
		fmt.Printf("The age is %v", age)
	} else {
		fmt.Println("Invalid name")
	}

	// iterate through maps - loops... can iterate through arrays also
	// iterating over map: no order is preserved


	// iterating through keys
	for name:= range myMap2 {
		fmt.Printf("Name: %v\n", name)
		
	}

	// iterating through key value pair
	for name, age := range myMap2 {
		fmt.Printf("Name: %v, Age: %v \n", name, age)
	}


	var intArr [3]int32 = [3]int32 {1,2,3}

	// iterating through array, index and value
	for i,v := range intArr {
		fmt.Printf("Index: %v, Value: %v \n", i, v)
	}

	// while loop in go

	var i int = 0

	for i <= 10{
		fmt.Println(i)
		i = i + 1
	}

	var j int = 0

	for {
		if j >= 10{
			break
		}
		fmt.Println(j)
		j = j + 1
	}

	for i:=0; i < 10; i++ { // can do all the other shorthands also like i--, i+= 10 etc
		fmt.Println(i)
	}

}