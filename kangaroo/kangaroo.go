package main

import "fmt"

func main() {
	arr := make([]int, 4)
	for i := 0; i < 4; i++ {
		fmt.Scan(&arr[i])
	}
	k1, k2 := 0, 0
	distance := getAbs((arr[0] + arr[1]) - (arr[2] + arr[3]))
	if distance == 0 {
		fmt.Println("YES")
		return
	}
	for i := 2; ; i++ {
		k1 = arr[0] + arr[1]*i
		k2 = arr[2] + arr[3]*i
		if k1 == k2 {
			fmt.Print("YES")
			return
		}
		if (getAbs(k1 - k2)) >= distance {
			fmt.Println("NO")
			return
		}
	}
}
func getAbs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
