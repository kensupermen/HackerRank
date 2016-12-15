package main

import (
	"fmt"
	"math/big"
)

func main() {
	var n int
	fmt.Scan(&n)
	result := big.NewInt(1)
	var i int64
	for i = 1; i <= int64(n); i++ {
		result.Mul(big.NewInt(i), result)
	}
	fmt.Println(result)
}
