package fibonacci

import "fmt"

func Print_fib_loop(n int) {
	a := 0
	b := 1
	for count := 1; count <= n; count++ {
		fmt.Printf("%d ", a)
		temp := a
		a = b
		b += temp
	}
	fmt.Println()
}

func Print_fib_recursion(n int) {
	for count := 1; count <= n; count++ {
		fmt.Printf("%d ", fib_nth(count))
	}
	fmt.Println()
}

func fib_nth(nth int) int {
	if nth == 1 {
		return 0
	} else if nth == 2 {
		return 1
	} else {
		return fib_nth(nth-1) + fib_nth(nth-2)
	}
}
