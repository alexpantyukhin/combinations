package combinations

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func getCombinationNumber(pr *Product) int {
	var number = 1

	for pr.Next() {
		number++
	}

	return number
}

func TestProduct_GeneratesCorrectNumberOfCombintations(t *testing.T) {
	pr2, _ := NewProduct([]interface{}{0, 1, 2}, 27)

	assert.Equal(t, 9, getCombinationNumber(pr2))
}
