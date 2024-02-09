package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(fmt.Sprintf("%v => %v", 101499, InsuranceRound(101499)))
	fmt.Println(fmt.Sprintf("%v => %v", 101500, InsuranceRound(101500)))
	fmt.Println(fmt.Sprintf("%v ===> %v", 101500, InsuranceRound(101500.01)))
	fmt.Println(fmt.Sprintf("%v ===> %v", 101500, InsuranceRound(101499.99)))
	fmt.Println(fmt.Sprintf("%v => %v", 101501, InsuranceRound(101501)))
	fmt.Println(fmt.Sprintf("%v => %v", 101501, InsuranceRound(101501)))
	fmt.Println(fmt.Sprintf("%v => %v", 799, InsuranceRound(599)))
	fmt.Println(fmt.Sprintf("%v => %v", 501, InsuranceRound(501)))
	fmt.Println(fmt.Sprintf("%v => %v", 500, InsuranceRound(500)))
	fmt.Println(fmt.Sprintf("%v => %v", 499, InsuranceRound(499)))
	fmt.Println(fmt.Sprintf("%v => %v", 100, InsuranceRound(100)))
	fmt.Println(fmt.Sprintf("%v => %v", 1000, InsuranceRound(1000)))
	fmt.Println(fmt.Sprintf("%v => %v", 10, InsuranceRound(10)))
	fmt.Println(math.Ceil(511))
}

func InsuranceRound(x float64) float64 {

	multiple := 500.0

	x = math.Ceil(x)

	return math.Ceil(x/multiple) * multiple
}
