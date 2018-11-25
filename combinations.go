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
	init          bool
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

	return &Product{currPosition, currValue, objs, repeat, false}, nil
}

// Next generates the next value for product
func (product *Product) Next() bool {
	if product.init == false {
		product.init = true
		return true
	}

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

// Permutation is
type Permutation struct {
	prod      *Product
	objs      []interface{}
	repeat    int
	currValue []interface{}
}

// NewPermutation is constructor
func NewPermutation(objs []interface{}, repeat int) (*Permutation, error) {
	pr, err := NewProduct(objs, repeat)

	if err != nil {
		return nil, err
	}

	return &Permutation{pr, objs, repeat, pr.Value()}, nil
}

// Next generates the next value for permutation
func (permutation *Permutation) Next() bool {
	pr := permutation.prod
	for pr.Next() {
		currValue := pr.Value()
		if checkArrayOnlyUniqueValues(pr.currPositions) {
			permutation.currValue = currValue
			return true
		}
	}

	return false
}

// Value gets the current value
func (permutation *Permutation) Value() []interface{} {
	return permutation.currValue
}

func checkArrayOnlyUniqueValues(indexes []int) bool {
	ln := len(indexes)
	for i := 0; i < ln-1; i++ {
		for j := i + 1; j < ln; j++ {
			if indexes[i] == indexes[j] {
				return false
			}
		}
	}

	return true
}
