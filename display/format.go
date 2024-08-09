package display

import (
	"fmt"
)

func printDemarcation() {
	demarcation := ""
	for i := 0; i < 15; i++ {
		demarcation += "-"
	}
	fmt.Println(demarcation)
}

func makeSameLength(x, y string) string {
	difference := len(x) - len(y)
	if difference > 0 {
		return x
	}

	difference *= -1
	result := ""

	for i := difference; i > 0; i-- {
		result += " "
	}
	return result + x
}

func printOperation(symbol string, initialValue int, fn func(x, y int) int, operands ...int) {
	result := initialValue
	for _, number := range operands {
		resultString := fmt.Sprintf("%d", result)
		numberString := symbol + " " + fmt.Sprintf("%d", number)
		fmt.Println(makeSameLength(resultString, numberString))
		fmt.Println(makeSameLength(numberString, resultString))
		printDemarcation()
		result = fn(result, number)

	}
}

func PrintAdd(fn func(x, y int) int, numbers ...int) {
	printOperation("+", 0, fn, numbers...)
}

func PrintMultiplication(fn func(x, y int) int, numbers ...int) {
	printOperation("*", 1, fn, numbers...)
}
