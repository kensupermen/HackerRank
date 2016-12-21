package main

import (
	"fmt"
	"strconv"
)

func main() {
	LIMIT := 1000001
	var n, m, tmp int
	fmt.Scan(&n)

	arrA := make([]int, LIMIT)
	for i := 0; i < n; i++ {
		fmt.Scan(&tmp)
		arrA[tmp]++
	}

	fmt.Scan(&m)

	arrB := make([]int, LIMIT)
	for i := 0; i < m; i++ {
		fmt.Scan(&tmp)
		arrB[tmp]++
	}

	for i := 0; i < LIMIT; i++ {
		if arrA[i] != arrB[i] {
			fmt.Print(strconv.Itoa(i) + " ")
		}
	}

}
