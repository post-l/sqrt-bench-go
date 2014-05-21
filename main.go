package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func Sqrt(n float64) float64 {
	if n == 0.0 {
		return 0.0
	}
	res := n / 2.0
	for i := 0; i < 10; i++ {
		res = 0.5 * (res + n/res)
	}
	return res
}

func main() {
	if len(os.Args) > 1 {
		in, err := strconv.ParseFloat(os.Args[1], 64)
		if err != nil {
			fmt.Printf("[Error] %v\n", err)
			return
		}
		out1 := Sqrt(in)
		out2 := math.Sqrt(in)
		fmt.Printf("Sqrt(%v) = %v\nmath.Sqrt(%v) = %v\n", in, out1, in, out2)
	} else {
		fmt.Printf("Usage: %s nbr\n", os.Args[0])
	}
}
