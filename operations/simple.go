package operations

import "github.com/kevmonstah/calculator/display"

func Add(numbers ...int) int {
	result := addThem(numbers...)
	return result
}

func addThem(numbers ...int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func Subtract(x, y int) int {
	result := subtractThem(x, y)
	return result
}

func subtractThem(x, y int) int {
	return x - y
}

func Multiply(numbers ...int) int {
	result := multiplyThem(numbers...)
	return result
}

func multiplyThem(numbers ...int) int {
	sum := 1
	for _, number := range numbers {
		sum *= number
	}
	return sum
}

func add(x, y int) int {
	return x + y
}

func multiply(x, y int) int {
	return x * y
}

func AdditionWithSteps(numbers ...int) {
	display.PrintAdd(add, numbers...)
}

func MultiplicationWithSteps(numbers ...int) {
	display.PrintMultiplication(multiply, numbers...)
}
