package combinations

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Enumerable interface {
	Next() bool
	Value() []interface{}
}

func getCombinationNumber(enum Enumerable) int {
	var number = 0

	for enum.Next() {
		number++
	}

	return number
}

func TestProduct_GeneratesCorrectNumberOfCombintations(t *testing.T) {
	pr, _ := NewProduct([]interface{}{0, 1, 2}, 3)

	assert.Equal(t, 27, getCombinationNumber(pr))
}

func TestPermutation_GeneratesCorrectNumberOfCombintationsFor3Objects3Repeats(t *testing.T) {
	per, _ := NewPermutation([]interface{}{0, 1, 2}, 3)

	assert.Equal(t, 6, getCombinationNumber(per))
}

func TestPermutation_GeneratesCorrectNumberOfCombintationsFor3Objects2Repeats(t *testing.T) {
	per, _ := NewPermutation([]interface{}{0, 1, 2}, 2)

	assert.Equal(t, 6, getCombinationNumber(per))
}

func TestPermutation_GeneratesCorrectNumberOfCombintationsFor3Objects1Repeats(t *testing.T) {
	per, _ := NewPermutation([]interface{}{0, 1, 2}, 1)

	assert.Equal(t, 3, getCombinationNumber(per))
}

func TestCombinations_GeneratesCorrectNumberOfCombintationsFor3Objects3Repeats(t *testing.T) {
	comb, _ := NewCombination([]interface{}{"A", "B", "C", "D"}, 2)

	assert.Equal(t, 6, getCombinationNumber(comb))
}

func TestCombinationsWithReplacement_GeneratesCorrectNumberOfCombintationsFor4Objects2Repeats(t *testing.T) {
	comb, _ := NewCombinationWithReplacement([]interface{}{"A", "B", "C", "D"}, 2)

	assert.Equal(t, 10, getCombinationNumber(comb))
}

func TestCombinationsWithReplacement_GeneratesCorrectNumberOfCombintationsFor4Objects1Repeats(t *testing.T) {
	comb, _ := NewCombinationWithReplacement([]interface{}{"A", "B", "C", "D"}, 1)

	assert.Equal(t, 4, getCombinationNumber(comb))
}
