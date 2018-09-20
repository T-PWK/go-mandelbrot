package mandelbrot

import (
	"testing"
)

func BenchmarkIteration(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Iter(0.1+0i, 20)
	}
}
func BenchmarkMandelbrot(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Mandelbrot(Min, Max, 20, 20, 20)
	}
}
