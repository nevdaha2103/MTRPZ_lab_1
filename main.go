package main

import (
	"math"
)

func computeRoots(a, b, c float64) (*float64, *float64) {
	discriminant := b*b - 4*a*c
	if discriminant < 0 {
		return nil, nil
	}
	if discriminant == 0 {
		root := -b / (2 * a)
		return &root, nil
	}
	sqrtVal := math.Sqrt(discriminant)
	root1 := (-b - sqrtVal) / (2 * a)
	root2 := (-b + sqrtVal) / (2 * a)
	return &root1, &root2
}
