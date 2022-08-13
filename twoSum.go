package main

import "fmt"

func main() {
	nums := []int{3, 2, 4}

	fmt.Println(twoSum(nums, 6))
}

func twoSum(nums []int, target int) []int {
	indexes := make(map[int]int)

	for idx, num := range nums {
		if found, ok := indexes[target-num]; ok {
			return []int{found, idx}
		}
		indexes[num] = idx
	}
	return nil
}
