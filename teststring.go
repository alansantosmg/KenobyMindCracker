package main

import (
	"errors"
	"fmt"
	"math"
	"strconv"
)

func main() {
	fmt.Println(IsNumeric("1"))
	fmt.Println(IsNumeric("12.345"))
	fmt.Println(IsNumeric("NOT"))
	x, xr := IsNumeric("1")
	if xr != nil {
		fmt.Println(xr)
	} else {
		fmt.Printf("%f", x)
	}

}

// IsNumeric é uma funçao
func IsNumeric(s string) (float64, error) {
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return float64(math.NaN()), errors.New("Não é numero")
	}

	return f, nil

}
