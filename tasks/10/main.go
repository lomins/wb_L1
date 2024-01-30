package main

import (
	"fmt"
	"math"
)

func main() {
	temps := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	groupedTemperatures := GroupTemperatures(temps, 10)

	for key, values := range groupedTemperatures {
		fmt.Printf("%d: %v\n", key, values)
	}

	fmt.Println()

	temps = []float64{0, 1, -10.4, -20.5, -35.6, 14}

	groupedTemperatures = GroupTemperatures(temps, 10)

	for key, values := range groupedTemperatures {
		fmt.Printf("%d: %v\n", key, values)
	}
}

func GroupTemperatures(temps []float64, step int) map[int][]float64 {
	groups := make(map[int][]float64)
	stepFloat := float64(step)

	for _, temp := range temps {
		groupKey := 0
		if temp > 0 {
			groupKey = int(math.Floor(temp/stepFloat)) * step
		} else {
			groupKey = int(math.Ceil(temp/stepFloat)) * step
		}
		groups[groupKey] = append(groups[groupKey], temp)
	}

	return groups
}
