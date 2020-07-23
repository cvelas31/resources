package main

import (
	"fmt"

	"gonum.org/v1/gonum/floats"
	"gonum.org/v1/gonum/stat"
)

func main() {
	x := make([]float64, 101)
	for i := range x {
		x[i] = 1.1 * float64(i) // x data ranges from 0 to 110
	}

	fmt.Println()
	fmt.Println("For ease, the floats Span function can be used to set the dividers")
	nBins := 1000
	dividers := make([]float64, nBins+1)
	min := floats.Min(x)
	max := floats.Max(x)
	// Increase the maximum divider so that the maximum value of x is contained
	// within the last bucket.
	fmt.Printf("dividers = %v\n", dividers)
	max++
	floats.Span(dividers, min, max)
	// Span includes the min and the max. Trim the dividers to create 10 buckets
	hist := stat.Histogram(nil, dividers, x, nil)
	fmt.Printf("Hist = %v\n", hist)
	// 	fmt.Println()
	// 	fmt.Println(`Histogram also works with weighted data, and allows reusing of
	// the count field in order to avoid extra garbage`)
	// 	weights := make([]float64, len(x))
	// 	for i := range weights {
	// 		weights[i] = float64(i + 1)
	// 	}
	// 	stat.Histogram(hist, dividers, x, weights)
	// 	fmt.Printf("Weighted Hist = %v\n", hist)
}
