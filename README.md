# Combinations
A package with combinations functions from [python itertools](https://docs.python.org/3/library/itertools.html).
For now the foolowing iterators are implemented :
   - [x] Product (the [Cartesian product](https://en.wikipedia.org/wiki/Cartesian_product)).
   - [x] Permutation (basing on the `Product` result).
   - [x] Combination.
   - [x] Combination with replacement.

# Installation
Just `go get` this repository with the following way:

```
go get github.com/alexpantyukhin/combinations
```

# Usage
```go
package main

import (
    "fmt"
    "github.com/alexpantyukhin/combinations"
)

func main() {
    inter := []interface{}{0, 1, 2}
    product, _ := combinations.NewProduct(inter, 3)

    for product.Next() {
        fmt.Println(product.Value())
    }
}
```
