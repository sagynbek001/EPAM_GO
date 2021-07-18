package main

import (
	"fmt"
	"sort"
	utf8 "unicode/utf8"
)

func main() {
	// Arrays task
	exampleArray := [10]int{1, 2, 3, 4, 2, 5, 6, 7, 8, 9}
	fmt.Println("1. Arrays task")
	fmt.Println("The given array is:", exampleArray)
	fmt.Printf("The average of the array is: %f\n\n", avgArray(exampleArray))

	// Slices task1
	exampleSlice1 := []string{"one", "two", "three", "four"}
	fmt.Println("2. Slices task1")
	fmt.Println("The given slice is:", exampleSlice1)
	fmt.Printf("The longest word in the slice is: %s\n\n", max(exampleSlice1))

	// Slices task2
	exampleSlice2 := []int64{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println("3. Slices task2")
	fmt.Println("The given slice is:", exampleSlice2)
	fmt.Println("The reversed slice is:", reverse(exampleSlice2))
	fmt.Println()

	// Maps task
	exampleMap := map[int]string{5: "one", 3: "two", 6: "three"}
	fmt.Println("4. Maps task1")
	fmt.Println("The given map is:", exampleMap)
	fmt.Println("The map values sorted in order of increasing keys:")
	printSorted(exampleMap)
}

func avgArray(arr [10]int) float64 {
	var total float64 = 0
	for _, v := range arr {
		total += float64(v)
	}
	return float64(total / 10)
}

func max(s []string) string {
	if len(s) > 0 {
		maxlen, answerIndex := utf8.RuneCountInString(s[0]), 0
		for i, v := range s {
			if utf8.RuneCountInString(v) > maxlen {
				maxlen, answerIndex = utf8.RuneCountInString(v), i
			}
		}
		return s[answerIndex]
	} else {
		return ""
	}
}

func reverse(s []int64) []int64 {
	newslice := []int64{}
	for i := len(s) - 1; i >= 0; i-- {
		newslice = append(newslice, s[i])
	}
	return newslice
}

func printSorted(m map[int]string) {
	keys := []int{}
	for i := range m {
		keys = append(keys, i)
	}
	sort.Ints(keys)
	for _, v := range keys {
		fmt.Print(m[v], " ")
	}
}
