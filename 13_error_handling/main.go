/* package main

import (
	"errors"
	"fmt"
)

func main() {
	result, err := eventNum(10)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Girdiğiniz sayı: ", result)
	}
}

func eventNum(num int) (int, error) {
	if num%2 != 0 {
		return 0, errors.New("HATA: Çift sayı girmediniz")
	}
	return num, nil
}
*/

/* package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {
	result, err := sRoot(-4)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

}

func sRoot(num float64) (float64, error) {
	if num < 0 {
		return 0, errors.New("Negatif sayıların karekökü alınamaz")
	}
	return math.Sqrt(num), nil

}
*/

package main

import (
	"fmt"
	"os"
)

func main() {

	file, err := os.Open("test.txt")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Dosyamız", file)
	}

}
