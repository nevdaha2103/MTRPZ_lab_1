package main

import (
	"bufio"
	"fmt"
	"math"
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

// Function to compute roots of a quadratic equation
func computeRoots(a, b, c float64) (*float64, *float64) {
	discriminant := b*b - 4*a*c
	if discriminant < 0 {
		// No real roots
		return nil, nil
	} else if discriminant == 0 {
		// One real root
		root := -b / (2 * a)
		return &root, nil
	} else {
		// Two real roots
		sqrtDisc := math.Sqrt(discriminant)
		root1 := (-b + sqrtDisc) / (2 * a)
		root2 := (-b - sqrtDisc) / (2 * a)
		return &root1, &root2
	}
}

func displayResults(root1, root2 *float64) {
	if root1 == nil {
		fmt.Println("No real roots found")
		return
	}
	if root2 == nil {
		fmt.Printf("One root found: %f\n", *root1)
		return
	}
	fmt.Printf("Two roots found: %f, %f\n", *root1, *root2)
}

func main() {
	var root1, root2 *float64
	if len(os.Args) <= 1 {
		a, b, c := getUserInput()
		root1, root2 = computeRoots(a, b, c)
	} else {
		a, b, c := fetchParamsFromFile()
		root1, root2 = computeRoots(a, b, c)
	}
	displayResults(root1, root2)
}
