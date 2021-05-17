/*
--- COMPOSITE TYPES
Slice - Struct - Pointer - Function - Interface - Map - Channel - Array
*/

// Package clause
package main

// Import statement
import (
	"fmt"
)

// My code
func main() {
	/*
		var (
			name      string  = "Fatih"
			age       uint8   = 40
			isMarried bool    = true
			weight    float32 = 72.5
		) */

	/* 	var name, age, isMarried, weight = "Fatih", 40, true, 72.5 */

	/* name, age, isMarried, weight := "Fatih", 40, true, 72.5 */

	/* var name string = "Fatih" */

	var name bool
	var age int

	fmt.Println(name, age)

}
