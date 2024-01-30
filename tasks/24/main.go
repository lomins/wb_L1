package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

func NewPoint(x, y float64) Point {
	return Point{x, y}
}

func FindDistance(p1 Point, p2 Point) float64 {
	dX := math.Abs(p1.x - p2.x)
	dY := math.Abs(p1.y - p2.y)

	return math.Sqrt(float64(dX*dX + dY*dY))
}

func main() {
	p1 := NewPoint(2, 4)
	p2 := NewPoint(7, 1)

	fmt.Printf("Distance: %.2f", FindDistance(p1, p2))
}
