package main

import "fmt"

func main() {
	var t int
	var str string
	count := 0
	fmt.Scan(&t)

	for i := 0; i < t; i++ {
		fmt.Scan(&str)
		count = 0
		arr1 := make([]int, 26)
		arr2 := make([]int, 26)
		j := len(str) - 1

		if len(str)%2 != 0 {
			count = -1
		} else {

			for i := 0; i < len(str)/2; i++ {
				arr1[str[i]-'a']++
				arr2[str[j]-'a']++
				j--
			}
			for i := 0; i < 26; i++ {
				if arr1[i] > arr2[i] {
					count += arr1[i] - arr2[i]
				}
			}

		}
		fmt.Println(count)
	}
}
