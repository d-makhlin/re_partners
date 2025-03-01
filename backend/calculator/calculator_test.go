package calculator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculatorSuccess(t *testing.T) {
	// test happy case
	c := NewCalculator([]int{})

	sizes, err := c.CalculatePacks(0)
	assert.NoError(t, err)
	assert.Equal(t, sizes, map[int]int{})
}

func TestSetNegativeSizes(t *testing.T) {
	// test cannot set negative sizes

	c := NewCalculator([]int{})

	err := c.SetPackSizes([]int{-1, 100, 200})
	assert.EqualError(t, err, "")
}

func TestFilterRepeatedPackSizes(t *testing.T) {
	// test repeated pack sizes are filtered

	c := NewCalculator([]int{})

	err := c.SetPackSizes([]int{100, 100, 200})
	assert.NoError(t, err)

	sizes, err := c.GetPackSizes()
	assert.NoError(t, err)

	assert.Equal(t, sizes, []int{100, 200})
}

func TestCheckMinAmount(t *testing.T) {
	// test of the second rule

	c := NewCalculator([]int{300, 500})

	sizesMap, err := c.CalculatePacks(400)
	assert.NoError(t, err)

	assert.Equal(t, sizesMap, map[int]int{500: 1})
}

func TestCheckMinPacks(t *testing.T) {
	// test of the third rule

	c := NewCalculator([]int{250, 300, 500})

	sizesMap, err := c.CalculatePacks(400)
	assert.NoError(t, err)

	assert.Equal(t, sizesMap, map[int]int{500: 1})
}

func TestCheckRulePrecedence(t *testing.T) {
	// test of the second rule precedence

	c := NewCalculator([]int{240, 300, 500})

	sizesMap, err := c.CalculatePacks(400)
	assert.NoError(t, err)

	assert.Equal(t, sizesMap, map[int]int{240: 2})
}

func TestNoPacks(t *testing.T) {
	// test the no packs case

	c := NewCalculator([]int{200, 300, 500})

	_, err := c.CalculatePacks(0)
	assert.EqualError(t, err, "")
}
