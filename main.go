package main

import (
	"errors"
	"fmt"
	"math"
)

func sqrt(x float64) (float64, error) {
	if x < 0 {
		return x, errors.New("not have negative number")
	}
	return math.Sqrt(x), nil
}

func main() {
	fmt.Println(sqrt(2))
	fmt.Println(sqrt(-2))
}
