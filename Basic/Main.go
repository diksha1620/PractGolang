package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	write := bufio.NewReader(os.Stdin)
	fmt.Println("Eneter your name : ")

	input, _ := write.ReadString('\n')
	fmt.Println("Customer name is:", input)

	f := []int{1, 2, 3, 66}

	f = append(f, 33)

	fmt.Println(f)
}
