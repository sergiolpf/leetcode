package main

import "fmt"

func maxProfit(prices []int) int {
	bestProfit := 0
	for i, value := range prices {
		max := findMax(prices[i:])
		if max-value > bestProfit {
			bestProfit = max - value
		}
	}

	return bestProfit
}

/*
	 func findMin(prices []int) (int, int) {
		min := prices[0]
		index := 0

		for i, value := range prices {
			if value < min {
				min = value
				index = i
			}
		}
		return index, min
	}
*/
func findMax(prices []int) int {
	max := prices[0]

	for _, value := range prices {
		if value > max {
			max = value
		}
	}

	return max
}

func main() {
	fmt.Printf("max profit is: %v", maxProfit([]int{}))
}
