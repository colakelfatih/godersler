package main

import "fmt"

/* var packVar = "Pack Scope" */

func main() {

	/* 	if true {
		var blockVar = "Block Scope"
		fmt.Println(blockVar)
	} */

	/* var funcVar = "Func Scope"
	fmt.Println(funcVar)
	fmt.Println(packVar)
	myFunc() */
	var name = "Fatih"
	name, surname := "Ahmet", "Colakel"
	fmt.Println(name, surname)

}

/* func myFunc() {
	fmt.Println(packVar)
} */
