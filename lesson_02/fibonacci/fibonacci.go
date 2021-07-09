package fibonacci

import "fmt"

func printFibLoop(n int) {
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

func printFibRecursion(n int) {
	for count := 1; count <= n; count++ {
		fmt.Printf("%d ", fibNth(count))
	}
	fmt.Println()
}

func fibNth(nth int) int {
	if nth == 1 {
		return 0
	} else if nth == 2 {
		return 1
	} else {
		return fibNth(nth-1) + fibNth(nth-2)
	}
}
