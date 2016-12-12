package main

import "fmt"

func main() {

	times := 5
  temp := make([]int64, times)
	var sum, max, min int64
	sum, max, min = 0,0,100000000000

	for i := 0; i < times; i++ {
		temp[i] = readStdin()
		sum += int64(temp[i])
		min = getMin(min, temp[i])
		max = getMax(max, temp[i])
	}
	fmt.Print(sum - int64(max), " ")
	fmt.Println(sum - int64(min))
}

func getMin(a, b int64) int64 {
	if (a > b) {
		return b
	}
  return a
}

func getMax(a, b int64) int64 {
	if (a > b) {
		return a
	}
	return b;
}

func readStdin() int64 {
	var n int64
	_, err := fmt.Scan(&n)
	if err != nil {
		panic(err)
	}
	return n
}
