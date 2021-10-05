package challenge

type flt func(int) bool

func IsOdd(n int) bool {
	return n % 2 != 0
}

func IsEven(n int) bool {
	return !IsOdd(n)
}

func Filter(sl[] int, f flt) (res[] int) {
	for _, num := range sl {
		if(f(num)) {
			res = append(res, num)
		} else {
			continue;
		}
	}
	return
}
