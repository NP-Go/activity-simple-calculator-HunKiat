package main

import (
	"fmt"
)

func add(a, b int) int {
	//Insert code here
	return a + b
}

func subtract(a, b int) int {
	//Insert code here
	return a - b
}

func multiply(a, b int) int {
	//Insert code here
	return a * b
}

func divide(a, b int) int {
	//Insert code here
	//consider for b = 0
	if b > 0 {
		return a / b // returns the quotient only
	}
	return 0
}

func main() {
	var a, b int
	var process string
	fmt.Println("Enter an integer: ")
	fmt.Scanln(&a)
	fmt.Println("Enter process: (add, sub, mul, div)")
	fmt.Scanln(&process)
	fmt.Println("Enter an integer: ")
	fmt.Scanln(&b)

	switch process {
	case "add":
		fmt.Printf("Add %d to %d gives %d", a, b, add(a, b))
	case "sub":
		fmt.Printf("Subtract %d from %d gives %d", b, a, subtract(a, b))
	case "mul":
		fmt.Printf("Multiply %d to %d gives %d", a, b, multiply(a, b))
	case "div":
		fmt.Printf("Divide %d over %d gives %d", a, b, divide(a, b))
	default:
		return
	}
	fmt.Println("")
	main()
}
