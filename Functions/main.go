package main

import "fmt"

func main() {

	for i := 0; i < 3; i++ {
		reader2("heloo", i)
	}

	fmt.Println("I am main Function")

	reader()
	var res1, res2, res3, res4 = add(11, 12)
	fmt.Println("addd: ", res1)
	fmt.Println("mul: ", res2)
	fmt.Println("devide: ", res3)
	fmt.Println("rem: ", res4)

	fmt.Println(moreadd(11, 12, 12, 33, 11))
	ans, err := devide(4, 0)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println(ans)
	}

}

// Error handling + Methods
func devide(a, b float64) (float64, error) {
	if b == 0.0 {
		// panic("not divisible by 0")
		return 0.0, fmt.Errorf("not divisible by 0")
	}
	return a / b, nil

}

func reader2(msg string, ind int) {
	fmt.Println(msg)
	fmt.Println(ind)
}
func reader() {
	fmt.Println("I am reader function")

}

func add(num1 int, num2 int) (int, int, int, int) {
	fmt.Println("Playing with 2 numbers")
	return num1 + num2, num1 * num2, num1 / num2, num1 % num2
}

func moreadd(numbers ...int) (total int) {

	fmt.Println("addition of numbers")
	for _, val := range numbers {
		total += val
	}
	return

}
