package main

import "fmt"

func main() {
	odd := []int{11, 33, 111, 66, 22}

	// without using range keyword
	for i := 0; i < len(odd); i++ {
		fmt.Printf("\nodd[%d] = %d \n", i, odd[i])
	}

	fmt.Printf("********************************")
	// Using range keyword
	for i, ele := range odd {
		fmt.Printf("\nodd[%d]= %d \n", i, ele)
	}

	fmt.Printf("********************************")
	// using range with string DS
	var str string = "DikshaBhutekar"
	for i, newStr := range str {
		fmt.Printf("\nstr[%d]= %v \n", i, string(newStr))
	}

	fmt.Printf("********************************")
	// createMap := make(map[string]int)
	createMap := map[string]int{
		"Diksha": 12,
		"Poonam": 13,
		"Ishhuu": 111,
		"Rushi":  124,
	}
	for newMap := range createMap {
		fmt.Printf("\nkeyvalue: %v and value: %d \n ", newMap, createMap[newMap])
	}

}
