package main

import "fmt"

func twoSum(nums []int, target int) []int {
	var indice1, indice2 int

	for index, value := range nums {
		for index2, value2 := range nums {
			if value+value2 == target && index < index2 {
				indice1 = index
				indice2 = index2
				break
			}
		}
	}
	return []int{indice1, indice2}
}
func main() {
	fmt.Println(twoSum([]int{}, 6))
}
