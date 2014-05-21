package main

import (
	"math"
	"testing"
)

func TestSqrt(t *testing.T) {
	const in, out = 4, 2

	if x := Sqrt(in); x != out {
		t.Errorf("Sqrt(%v) = %v, want %v", in, x, out)
	}
}

func TestSqrt2(t *testing.T) {
	for in := 0; in < 1000; in++ {
		out := float64(int(math.Sqrt(float64(in))*100)) / 100
		x := float64(int(Sqrt(float64(in))*100)) / 100
		if x != out {
			t.Errorf("Sqrt(%v) = %v, want %v", in, x, out)
		}
	}
}

func BenchmarkSqrt(b *testing.B) {
	for in := 0; in < b.N; in++ {
		Sqrt(float64(in))
	}
}

func BenchmarkMathSqrt(b *testing.B) {
	for in := 0; in < b.N; in++ {
		math.Sqrt(float64(in))
	}
}
