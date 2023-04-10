package main

import "fmt"

func pivotIndex(nums []int) int {

	for i := 0; i < len(nums); i++ {
		leftSum := getSum(nums[0:i])
		rightSum := getSum(nums[i+1:])

		if leftSum == rightSum {
			return i
		}
	}

	return -1
}

func getSum(nums []int) int {
	sum := 0

	for _, value := range nums {
		sum += value
	}

	return sum
}

// this is definitely O(n)
func pivotIndex2(nums []int) int {
	leftSum := []int{}
	rightSum := make([]int, len(nums))

	//pivot := -1
	sum := 0

	for _, v := range nums {
		sum += v
		leftSum = append(leftSum, sum)
	}
	sum = 0

	for i := len(nums) - 1; i >= 0; i-- {
		sum += nums[i]
		rightSum[i] = sum
	}

	for i := 0; i < len(nums); i++ {
		if leftSum[i] == rightSum[i] {
			return i
		}
	}

	return -1

}

func pivotIndex3(nums []int) int {
	totalSum := 0

	for _, num := range nums {
		totalSum += num
	}

	leftSum := 0

	for i, num := range nums {
		if leftSum == totalSum-leftSum-num {
			return i
		}
		leftSum += num
	}

	return -1
}
func main() {
	fmt.Println(pivotIndex([]int{}))
}
