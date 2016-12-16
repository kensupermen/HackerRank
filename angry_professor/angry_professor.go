package main

import "fmt"

func main() {
	var testcase, k, n int
	fmt.Scan(&testcase)
	result := make([]string, testcase)
	for i := 0; i < testcase; i++ {
		result[i] = "YES"
		fmt.Scan(&n, &k)
		countOnTime := 0
		for j := 0; j < n; j++ {

			var val int
			fmt.Scan(&val)
			if val <= 0 {
				countOnTime++
			}
			if countOnTime >= k {
				result[i] = "NO"
				// break
			}
		}
	}
	for i := 0; i < testcase; i++ {
		fmt.Println(result[i])
	}
}
