package main

import (
	"fmt"
)

func main() {
	a := make(map[string]int)
	fmt.Println(a)
	a = map[string]int{
		"Happy":    11,
		"Hppiesr":  12,
		"Happiest": 13,
	}
	fmt.Println(a)

	a["save"] = 22
	fmt.Println(a)

	pop, err := a["save"]
	delete(a, "Happpiest")
	fmt.Println(pop, err)

}
