package main

import (
	"fmt"
)

func CountBy(x, n int) []int {
	res := make([]int, n)
	for i := 1; i <= n; i++ {
		res[i-1] = x * i
	}
	return res
}

func main() {
	// numbers := []int{1, 10}

	fmt.Println(CountBy(1, 10))

}
