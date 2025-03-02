package calculator

import (
	"fmt"
	"slices"
)

type Calculator struct {
	packSizes []int
}

func NewCalculator(packSizes []int) *Calculator {
	// Init and return a new calculator instance

	// ensure the sizes are sorted and have no duplicates
	slices.Sort(packSizes)
	return &Calculator{packSizes: slices.Compact(packSizes)}
}

func (c *Calculator) SetPackSizes(packSizes []int) error {
	// Update the packs' sizes

	if len(packSizes) == 0 {
		return fmt.Errorf("You should provide pack sizes")
	}
	sizes := []int{}
	for _, elem := range packSizes {
		if elem <= 0 {
			return fmt.Errorf("all sizes should be positive, %d is not", elem)
		}
		sizes = append(sizes, elem)
	}

	// ensure the sizes are sorted and have no duplicates
	slices.Sort(sizes)
	c.packSizes = slices.Compact(sizes)
	return nil
}

func (c *Calculator) GetPackSizes() ([]int, error) {
	return c.packSizes, nil
}

func copyMap(original map[int]int) map[int]int {
	copy := make(map[int]int)
	for k, v := range original {
		copy[k] = v
	}
	return copy
}

func (c *Calculator) CalculatePacks(amount int) (map[int]int, error) {
	// Calculate the necessary packs collection

	// 1. Only whole packs can be sent. Packs cannot be broken open.
	// 2. Within the constraints of Rule 1 above, send out the least amount of items to fulfil the order.
	// 3. Within the constraints of Rules 1 & 2 above, send out as few packs as possible to fulfil each order.
	// (Please note, rule #2 takes precedence over rule #3)

	// Use dynamic programming, every dp array index is a sum
	// every value related to the index - minimum packs required to reach the sum

	// ensure that we can store the sum in case it is equal to the amount + the largest pack size
	if amount <= 0 {
		return nil, fmt.Errorf("The amount should be positive, %d is not", amount)
	}
	dp := make([]int, amount+c.packSizes[len(c.packSizes)-1]+1)

	// store the minimum map of the used packs for every reached sum
	count := make([]map[int]int, amount+c.packSizes[len(c.packSizes)-1]+1)

	for i := range dp {
		dp[i] = amount + 1
	}
	dp[0] = 0
	for i := 0; i <= amount; i++ {
		for _, num := range c.packSizes {
			if dp[i]+1 < dp[i+num] { // If using 'num' gives us fewer elements for sum i+num
				dp[i+num] = dp[i] + 1
				count[i+num] = copyMap(count[i]) // Copy the map of used packs in since there are fewer elements used
				count[i+num][num]++              // Increment the count for the used pack
			}
		}
	}

	// find the closest sum to the required amount
	for i := amount; i < amount+c.packSizes[len(c.packSizes)-1]; i++ {
		if dp[i] < amount+1 {
			return count[i], nil
		}
	}
	return nil, fmt.Errorf("the sum was not found")

}
