package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	arrInput := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arrInput[i])
	}
	// fmt.Println("---------------------")
	for i := 0; i < n; i++ {
		fmt.Println(calc(arrInput[i]))
	}
}

func calc(t int) int {
	result := 1 // origin
	if t >= 1 { // spring
		result++ // = 1 * 2
	}
	for i := 2; i < t; i += 2 { // each year
		result++            // summer
		result = result * 2 // spring
	}

	if t > 1 && t%2 == 0 {
		result++ // summer
	}
	return result
}
