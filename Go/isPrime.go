package main

func modMultiply(a, b, mod uint64) uint64 {
	res := uint64(0)
	a = a % mod
	for b > 0 {
		if b%2 == 1 {
			res = (res + a) % mod
		}
		a = (2 * a) % mod
		b /= 2
	}
	return res % mod
}

func binPower(base, e, mod uint64) uint64 {
	// fmt.Println("BinP", base, e, mod)
	var result uint64 = 1
	base %= mod
	for e > 0 {
		if (e % 2) == 1 {
			result = modMultiply(result, base, mod)
		}
		base = modMultiply(base, base, mod)
		// fmt.Println("Base : ", base, e)
		e = e / 2
	}
	return result
}

func checkComposite(n, a, d uint64, s int) bool {
	// fmt.Println("E", n, a, d, s)
	var x uint64 = binPower(a, d, n)

	// fmt.Println("a", x)
	if (x == 1) || (x == n-1) {
		return false
	}
	for r := 1; r < s; r++ {
		x = modMultiply(x, x, n)
		if x == n-1 {
			// fmt.Println("a", x)
			return false
		}
	}

	return true
}

func MillerRabin(n uint64) bool {
	if n < 2 {
		return false
	}
	r := 0
	var d uint64 = n - 1

	for (d & 1) == 0 {
		d = d / 2
		r++
	}

	a := [12]uint64{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37}

	for _, val := range a {

		if n == val {
			// fmt.Printf("%v 0\n", n)
			return true

		}
		if checkComposite(n, val, d, r) {
			// fmt.Println("HAHA", n, val, d, r)
			// fmt.Printf("%v 0\n", n)
			return false
		}
	}

	// fmt.Printf("%v 1\n", n)
	return true
}
