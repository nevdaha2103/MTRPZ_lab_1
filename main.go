package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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

func fetchParamsFromFile() (float64, float64, float64) {
	var coefficients [3]float64
	filePath := os.Args[1]
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var content string
	if scanner.Scan() {
		content = scanner.Text()
	}
	params := strings.Split(content, " ")
	if len(params) != 3 {
		panic("invalid input")
	}
	for i := 0; i < 3; i++ {
		coefficients[i], err = parseFloat(params[i])
		if err != nil {
			panic(err)
		}
	}
	return coefficients[0], coefficients[1], coefficients[2]
}
