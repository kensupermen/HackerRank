package main

import "fmt"

func main() {

	var t, m, n, tmp int
	fmt.Scan(&t)
	for j := 0; j < t; j++ {
		fmt.Scan(&m)
		fmt.Scan(&n)
		arrResult := make([]int, 10001)
		arrInput := make([]int, n+1)
		for i := 1; i <= n; i++ {
			fmt.Scan(&tmp)
			if tmp >= m {
				continue
			}
			arrInput[i] = tmp
			arrResult[tmp] = i
		}

		for i := 1; i <= n; i++ {
			if arrResult[m-arrInput[i]] > 0 {
				fmt.Println(i, arrResult[m-arrInput[i]])
				break
			}
		}

	}
}
