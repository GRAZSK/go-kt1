package main

import (
	"fmt"
	"math"
	"sort"
)

func groupByStep(temps []float64, step float64) map[float64][]float64 {
	groups := make(map[float64][]float64) 

	for _, t := range temps {
	
		key := math.Round(t/step) * step
		groups[key] = append(groups[key], t)
	}

	return groups
}

func main() {
	temps := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	groups := groupByStep(temps, 10)
	keys := make([]float64, 0, len(groups))
	for k := range groups {
		keys = append(keys, k)
	}
	sort.Float64s(keys)

	for _, k := range keys {
		fmt.Printf("%.0f: %v\n", k, groups[k])
	}
}
