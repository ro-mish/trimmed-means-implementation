package main

import (
	"fmt"
	"math/rand"

	trimmedmean "trimmed-mean"
)

func generateIntData(n int) []int {
	// Initialize empty slice
	data := make([]int, n)

	// Generate random integers
	for i := range data {
		data[i] = rand.Intn(1000)
	}
	return data
}

func generateFloatData(n int) []float64 {
	// Initialize empty slice
	data := make([]float64, n)
	for i := range data {
		data[i] = rand.Float64() * 1000
	}
	return data
}

func main() {
	// Set random seed once at program start
	rand.Seed(1234)

	const n = 100
	const trimPct = 0.05

	numTrimmed := int(float64(n) * trimPct)

	fmt.Printf("Sample size: %d\n", n)
	fmt.Printf("Trim percentage: %.1f%%\n", trimPct*100)
	fmt.Printf("Num values trimmed from each end: %d\n", numTrimmed)
	fmt.Printf("Num values used in mean: %d\n\n", n-2*numTrimmed)

	// Test with integers
	intData := generateIntData(n)
	intResult, err := trimmedmean.TrimmedMeanInt(intData, trimPct)
	if err != nil {
		fmt.Printf("Error calculating trimmed mean for integers: %v\n", err)
	} else {
		fmt.Printf("Integer Data head(10): %v\n", intData[:10])
		fmt.Printf("Integer %g%% Trimmed Mean: %.2f\n\n", trimPct*100, intResult)
	}

	// Test with floats
	floatData := generateFloatData(n)
	floatResult, err := trimmedmean.TrimmedMean(floatData, trimPct)
	if err != nil {
		fmt.Printf("Error calculating trimmed mean for floats: %v\n", err)
	} else {
		fmt.Printf("Float Data head(10): %.2f\n", floatData[:10])
		fmt.Printf("Float %g%% Trimmed Mean: %.2f\n\n", trimPct*100, floatResult)
	}
}
