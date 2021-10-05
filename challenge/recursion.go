package challenge

func Factorial(n uint64) (fac uint64) {
	if(n == 1) {
		fac = n
	} else {
		fac = n * Factorial(n-1)
	}
	return
}