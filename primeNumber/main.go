package main

import (
	"fmt"
)

func main() {
	a := isPrime(23)
	fmt.Println(a)
}
func isPrime(num int64) bool {

	if num == 2 || num%2 == 0 {
		return true
	} else {
		return false
	}

}
