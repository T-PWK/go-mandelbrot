package mandelbrot

const (
	// Min value for a default Mandelbrot fractal
	Min = -2 + 1.4i

	// Max value for a default Mandelbrot fractal
	Max = 0.8 - 1.4i
)

// Mandelbrot returns an image of the Mandelbrot fractal of size (h,w)
// between lu (left-upper) and rb (right-bottom) complex numbers with a specifed maximum number of iterations
func Mandelbrot(lu, rb complex128, h, w, maxit uint) [][]uint {

	if h < 2 || w < 2 {
		panic("Invalid image size")
	}

	if real(lu) > real(rb) || imag(rb) > imag(lu) {
		panic("Invalid complex numbers boundries")
	}

	rmin := real(lu)
	imin := imag(rb)
	dx := (real(rb) - real(lu)) / float64(w-1)
	dy := (imag(lu) - imag(rb)) / float64(h-1)

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

// Iter calculates iteration value for a given complex number (point)
func Iter(p complex128, maxit uint) uint {
	z := p

	for it := uint(0); it < maxit; it++ {
		z = z*z + p

		if r, i := real(z), imag(z); r*r+i*i > 4 {
			return it
		}
	}

	return maxit
}
