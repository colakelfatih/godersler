/* package main

import "fmt"

func main() {

	merhaba("Fatih", 26) // Argument Function
}

func merhaba(name string, age int) { // Parametre Function
	fmt.Printf("Benim adım %s, yaşım %d", name, age)
}
*/

/* package main

import "fmt"

func main() {
	fmt.Println(result(55))
}

func result(grade int) string {
	if grade >= 50 {
		return "geçtiniz"
	}
	return "kaldınız"

}
*/

/*  package main

import "fmt"

func main() {
	fmt.Println(result(55))
}

func result(grade int) string {
	if grade >= 50 {
		return "geçtiniz"
	}
	return "kaldınız"

} */

/* package main

import (
	"fmt"
	"time"
)

func main() {
	var today time.Time = time.Now() // method
	fmt.Println(today)
} */

/* package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Print("Lütfen Sınav Sonucunu Giriniz = ")
	reader := bufio.NewReader(os.Stdin)
	value, _ := reader.ReadString('\n') // blank identifier
	fmt.Println(value)
} */

package main

import "fmt"

func main() {
	bolum, kalan := bolme(104, 5)
	fmt.Println(bolum, kalan)

}

//104 / 5 ==> 20 - 4
func bolme(bolunen, bolen int) (bolum, kalan int) {
	bolum = bolunen / bolen
	kalan = bolunen % bolen

	return bolum, kalan

}
