package fibonacci_test

import (
	"strconv"
	"testing"

	"../fibonacci"
)

func TestCallAdd(t *testing.T) {
	result := fibonacci.CalAdd(1, 2)
	if result != 3 {
		t.Error("Expected 3, got", result)
	}
}

func TestFibonancy(t *testing.T) {
	_, _, fibonacci := fibonacci.Fibonancy(2, 3)
	if fibonacci != 5 {
		t.Error("Expected 5, got", strconv.Itoa(fibonacci))
	}
}
