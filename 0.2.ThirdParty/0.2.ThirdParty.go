// 0.2 Third Party
package main

import (
	"fmt"

	"gonum.org/v1/gonum/stat"
)

func main() {
	values := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(stat.Variance(values, nil))
}
