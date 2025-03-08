package main

import (
	"fmt"
	"strconv"
)

func parseFloat(input string) (float64, error) {
	value, err := strconv.ParseFloat(input, 64)
	if err != nil {
		return 0, fmt.Errorf("invalid number: %v", err)
	}
	return value, nil
}

func getUserInput() (float64, float64, float64) {
	var inputs [3]string
	var coefficients [3]float64
	var err error
	labels := [3]string{"a", "b", "c"}
	fmt.Println("Please enter the coefficients:")

	for i, label := range labels {
		for {
			fmt.Printf("%v = ", label)
			fmt.Scan(&inputs[i])
			coefficients[i], err = parseFloat(inputs[i])
			if err != nil {
				fmt.Printf("Error. Expected a valid real number, got '%v' instead\n", inputs[i])
				continue
			}
			break
		}
	}
	return coefficients[0], coefficients[1], coefficients[2]
}
