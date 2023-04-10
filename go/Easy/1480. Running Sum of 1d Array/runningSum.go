package main

import "fmt"

// normal approach, maybe O(n) space O(1)
func runningSum(nums []int) []int {
	var output []int
	var sum int = 0

	for _, value := range nums {
		sum += value

		output = append(output, sum)
	}

	return output
}

// optimized approach, maybe O(n) space O(1)
func runningSum2(nums []int) []int {
	for i := 1; i < len(nums); i++ {
		nums[i] += nums[i-1]
	}

	return nums
}

func main() {
	fmt.Println(runningSum([]int{}))
}
