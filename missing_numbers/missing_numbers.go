package main

import (
	"fmt"
	"sort"
)

func main() {
	var n, m int
	fmt.Scan(&n)

	arrA := make(map[int]int)
	for i := 0; i < n; i++ {
		var value int
		fmt.Scan(&value)
		arrA[value]++
	}

	fmt.Scan(&m)

	arrB := make(map[int]int)

	for i := 0; i < m; i++ {
		var value int
		fmt.Scan(&value)
		arrB[value]++
	}

	// fmt.Println("----------------Process----------------")

	for kA, vA := range arrA {
		for kB, vB := range arrB {
			if kA == kB {
				arrB[kB] = vB - vA
			}
		}
	}

	// fmt.Println("----------------Result----------------")
	keys := make([]int, 0, len(arrB))
	for k, v := range arrB {
		if v > 0 {
			keys = append(keys, k)
		}
	}
	sort.Ints(keys)
	for i := 0; i < len(keys); i++ {
		fmt.Print(keys[i], " ")
	}
}
