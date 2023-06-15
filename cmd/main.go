package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("CMD line ")
	argwp := os.Args
	argswop := os.Args[1:]
	arg := os.Args[3:]
	fmt.Println(argswop)
	fmt.Println(argwp)
	fmt.Println(arg)
}
