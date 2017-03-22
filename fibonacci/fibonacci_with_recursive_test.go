package fibonacci_test

import (
	"strings"
	"testing"

	"../fibonacci"
)

func TestFibonancyWithRecursive(t *testing.T) {
	result := fibonacci.FibonancyWithRecursive(0, 1, 8, "")
	var expectData = "0,1,1,2,3,5,8"
	if strings.Compare(result, expectData) != 0 {
		t.Error("Expected "+expectData+" got", result)
	}
}
