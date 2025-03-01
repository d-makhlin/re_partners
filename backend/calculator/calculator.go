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

func (c *Calculator) CalculatePacks(amount int) (map[int]int, error) {
	// Calculate the necessary packs collection

	// 1. Only whole packs can be sent. Packs cannot be broken open.
	// 2. Within the constraints of Rule 1 above, send out the least amount of items to fulfil the order.
	// 3. Within the constraints of Rules 1 & 2 above, send out as few packs as possible to fulfil each order.
	// (Please note, rule #2 takes precedence over rule #3)

	packsMap := map[int]int{}
	for _, elem := range c.packSizes {
		packsMap[elem] = 0
	}

	return packsMap, nil

}
