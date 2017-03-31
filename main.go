package main

import (
	"fmt"

	"./fibonacci"
)

func main() {
	// fmt.Println("Fibonacci With For-Loop")
	// fibonacci.LoopDisplay(0, 1, 5)
	// fmt.Println("Fibonacci With Recursive Function")
	fmt.Printf("%s\n", fibonacci.FibonancyWithRecursiveDefault(25))
}
