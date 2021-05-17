package main

import "fmt"

func main() {
	//fmt.Println(sum(5, 10))
	//merhaba()
	/* z := sum(5, 10)
	sum2(7, 6)
	fmt.Println(z) */
	merhaba2("fatih", 26)

}

func sum(x, y int) int {
	return x + y
}

func sum2(x, y int) {
	fmt.Println(x + y)
}

func merhaba() {
	fmt.Println("Benim adım Fatih")
}

func merhaba2(name string, age int) {
	fmt.Printf("Adım %s, yaşım %d", name, age)
}
