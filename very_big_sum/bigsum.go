package main

import "fmt"

func main() {
	var times, temp int
	var result int64

	// fmt.Println("How many number you want to sum?")
	times = readStdin()
	for i := 0; i < times; i++ {
		// fmt.Println("Input the number ", i)
		temp = readStdin()
		result += int64(temp)
	}

	fmt.Println(result)
}

func readStdin() int {
	var n int
	_, err := fmt.Scan(&n)
	if err != nil {
		panic(err)
	}
	return n
}
