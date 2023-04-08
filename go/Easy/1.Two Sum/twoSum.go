package main

import "fmt"

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)

	for index, value := range nums {
		complement := target - value
		num, hasFound := m[complement]
		if hasFound {
			return []int{num, index}
		}

		m[value] = index

	}
	return nil
}
func main() {
	fmt.Println(twoSum([]int{}, 6))
}
