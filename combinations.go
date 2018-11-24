package combinations

import (
	"errors"
)

// Product is carthesian product
type Product struct {
	currPositions []int
	currValue     []interface{}
	objs          []interface{}
	repeat        int
}

// NewProduct is constructor
func NewProduct(objs []interface{}, repeat int) (*Product, error) {
	if repeat < 1 {
		return nil, errors.New("len should be greater then 0")
	}

	if len(objs) < 1 {
		return nil, errors.New("length of objs should be greater then 0")
	}

	currPosition := make([]int, repeat)
	currValue := make([]interface{}, repeat)

	for i := 0; i < repeat; i++ {
		currPosition[i] = 0
		currValue[i] = objs[0]
	}

	return &Product{currPosition, currValue, objs, repeat}, nil
}

// Next generates the next value for product
func (product *Product) Next() bool {
	maxIndex := len(product.objs) - 1

	numberMaxIndexes := 0
	for i := product.repeat - 1; i >= 0; i-- {
		if product.currPositions[i] == maxIndex {
			numberMaxIndexes++
		}
	}

	if numberMaxIndexes == product.repeat {
		return false
	}

	for i := product.repeat - 1; i >= 0; i-- {
		if product.currPositions[i] < maxIndex {
			product.currPositions[i]++
			product.currValue[i] = product.objs[product.currPositions[i]]

			break
		}

		product.currPositions[i] = 0
		product.currValue[i] = product.objs[0]
	}

	return true
}

// Value gets the current value
func (product *Product) Value() []interface{} {
	return product.currValue
}
