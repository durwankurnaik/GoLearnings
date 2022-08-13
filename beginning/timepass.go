package main

import (
	"fmt"
	"sort"
)

// Algo-expert hard level question
func main() {
	arr := []int{4, 7, 14, 19, 32, 1, 3, 5, 2, 8, 9, 10, 11, 12, 13}
	lenArr := len(arr)
	count := 1

	sort.Ints(arr)

	ans := [2]int{arr[0], arr[0]}
	for i := 0; i < lenArr-1; i++ {
		j := i
		currCount := 1
		var tempAns [2]int
		tempAns[0] = arr[j]

		for j < lenArr-1 && arr[j]+1 == arr[j+1] {
			currCount++
			j++
		}

		tempAns[1] = arr[j]
		i = j

		if currCount > count {
			count = currCount
			ans = tempAns
		}
	}
	fmt.Println(ans)
}
