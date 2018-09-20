// Package mandelbrot provides functions for generating a Mandelbrot set.
package mandelbrot

const (
	// Min upper-left point for a default Mandelbrot fractal
	Min = -2 + 1.4i

	// Max lower-right point for a default Mandelbrot fractal
	Max = 0.8 - 1.4i
)

// Mandelbrot returns the Mandelbrot set of size (h,w).
// The set is generated between ul (upper-left) and lr (lower-right) points with the maximum number of iterations.
func Mandelbrot(ul, lr complex128, h, w, maxit uint) [][]uint {

	if h < 2 || w < 2 {
		panic("Invalid image size")
	}

	if real(ul) > real(lr) || imag(lr) > imag(ul) {
		panic("Invalid corner points")
	}

	rmin := real(ul)
	imin := imag(lr)
	dx := (real(lr) - real(ul)) / float64(w-1)
	dy := (imag(ul) - imag(lr)) / float64(h-1)

	c := make([][]uint, h)

	for i := range c {
		c[i] = make([]uint, w)
	}

	for y := uint(0); y < h; y++ {
		for x := uint(0); x < w; x++ {
			c[y][x] = Iter(complex(rmin+float64(x)*dx, imin+float64(y)*dy), maxit)
		}
	}

	return c
}

// Iter calculates iteration value for a given number point.
// The iteration value uses the following algorithm:
//   z = p;
//   repeat
//     z = z*z + p
//   until |z| < 2;
// It returns number of iteration at which the |z| < 2 condition is not met or the maximum iterations value.
func Iter(p complex128, maxit uint) uint {
	z := p

	for it := uint(0); it < maxit; it++ {
		z = z*z + p

		// For better performance we use r^2 + i^2 > 4 instead of cmplx.Abs(z) > 2
		if r, i := real(z), imag(z); r*r+i*i > 4 {
			return it
		}
	}

	return maxit
}
