package fibonacci

import "strconv"

func FibonancyWithRecursiveDefault(max int) string {
	return FibonancyWithRecursive(0, 1, max, "")
}

func FibonancyWithRecursive(digit1, digit2, max int, result string) string {
	if digit2 > max {
		result += strconv.Itoa(digit1)
		return result
	}
	result += strconv.Itoa(digit1) + ","
	return FibonancyWithRecursive(digit2, digit1+digit2, max, result)
}
