package fibonacci

import (
	"fmt"
	"strconv"
)

func CalAdd(number1, number2 int) int {
	return number1 + number2
}

func Fibonancy(digit1, digit2 int) (int, int, int) {
	var currentNum, fiboNum = digit2, CalAdd(digit1, digit2)
	return fiboNum - currentNum, currentNum, fiboNum
}

func Display(fibonancy int) {
	if fibonancy != 0 {
		fmt.Printf("%s", ",")
	}
	fmt.Printf("%s", strconv.Itoa(fibonancy))
}

func LoopDisplay(digit1, digit2 int, maxRound int) {
	for count := 0; count < maxRound; count++ {
		var a int
		a, digit1, digit2 = Fibonancy(digit1, digit2)
		if a == 0 && digit1 == 1 {
			Display(0)
			Display(1)
		}
		Display(digit2)
	}
	println()
}
