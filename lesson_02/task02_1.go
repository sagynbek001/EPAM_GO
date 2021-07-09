package main

import (
	"fmt"
	
	fib "github.com/sagynbek001/EPAM_GO/fibonacci"
)

func main() {
	defer fmt.Println("Thank you for running this program!")
	fmt.Println("Hello, this application prints first 10 Fibonacci numbers!")
	fmt.Println("Using a loop:")
	fib.Print_fib_loop(10)
	fmt.Println("Using recursion:")
	fib.Print_fib_recursion(10)
}
