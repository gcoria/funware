package main

import (
	"fmt"
)

// maxProfit returns the maximum profit from at most k stock transactions
func maxProfit(k int, prices []int) int {
	n := len(prices)
	if n <= 1 || k == 0 {
		return 0
	}

	// If k >= n/2, we can make as many transactions as we want
	// This becomes the unlimited transactions problem
	if k >= n/2 {
		return maxProfitUnlimited(prices)
	}

	// DP approach for limited transactions
	// buy[i] = max profit after at most i transactions, currently holding stock
	// sell[i] = max profit after at most i transactions, not holding stock
	buy := make([]int, k+1)
	sell := make([]int, k+1)

	// Initialize: buying on day 0
	for i := 0; i <= k; i++ {
		buy[i] = -prices[0]
		sell[i] = 0
	}

	// Process each day
	for i := 1; i < n; i++ {
		// Process transactions in reverse order to avoid using updated values
		for j := k; j >= 1; j-- {
			// sell[j] = max(keep previous sell[j], sell current stock)
			sell[j] = max(sell[j], buy[j]+prices[i])
			// buy[j] = max(keep previous buy[j], buy stock using prev transaction)
			buy[j] = max(buy[j], sell[j-1]-prices[i])
		}
	}

	return sell[k]
}

// maxProfitUnlimited handles the case where we can make unlimited transactions
func maxProfitUnlimited(prices []int) int {
	profit := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			profit += prices[i] - prices[i-1]
		}
	}
	return profit
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	// Test with the sample input
	prices := []int{5, 11, 3, 50, 60, 90}
	k := 2

	result := maxProfit(k, prices)
	fmt.Printf("Input: prices = %v, k = %d\n", prices, k)
	fmt.Printf("Output: %d\n", result)
	fmt.Printf("Explanation: Buy at 5, sell at 11 (profit: 6); Buy at 3, sell at 90 (profit: 87); Total: %d\n", result)

	// Additional test cases
	fmt.Println("\nAdditional test cases:")

	// Test case 1
	prices1 := []int{2, 4, 1}
	k1 := 2
	result1 := maxProfit(k1, prices1)
	fmt.Printf("prices = %v, k = %d => %d\n", prices1, k1, result1)

	// Test case 2
	prices2 := []int{3, 2, 6, 5, 0, 3}
	k2 := 2
	result2 := maxProfit(k2, prices2)
	fmt.Printf("prices = %v, k = %d => %d\n", prices2, k2, result2)
}
