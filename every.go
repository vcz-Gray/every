package every

type ints func(int) bool
type floats func(float64) bool
type strings func(string) bool
type bytes func(byte) bool
type bools func(bool) bool
type complexes func(complex128) bool
type anythings func(interface{}) bool

func Ints(n []int, f ints) bool {
	for _, el := range n {
		if f(el) == false {
			return false
		}
	}
	return true
}

func Floats(n []float64, f floats) bool {
	for _, el := range n {
		if f(el) == false {
			return false
		}
	}
	return true
}

func Strings(s []string, f strings) bool {
	for _, el := range s {
		if f(el) == false {
			return false
		}
	}
	return true
}

func Bytes(s []byte, f bytes) bool {
	for _, el := range s {
		if f(el) == false {
			return false
		}
	}
	return true
}

func Bools(b []bool, f bools) bool {
	for _, el := range b {
		if f(el) == false {
			return false
		}
	}
	return true
}

func Complexes(c []complex128, f complexes) bool {
	for _, el := range c {
		if f(el) == false {
			return false
		}
	}
	return true
}

func Any(a interface{}, f anythings) bool {
	if n, ok := a.([]int); ok {
		for _, el := range n {
			if f(el) == false {
				return false
			}
		}
		return true
	}
	if n, ok := a.([]float64); ok {
		for _, el := range n {
			if f(el) == false {
				return false
			}
		}
		return true
	}
	if n, ok := a.([]string); ok {
		for _, el := range n {
			if f(el) == false {
				return false
			}
		}
		return true
	}
	if n, ok := a.([]byte); ok {
		for _, el := range n {
			if f(el) == false {
				return false
			}
		}
		return true
	}
	if n, ok := a.([]bool); ok {
		for _, el := range n {
			if f(el) == false {
				return false
			}
		}
		return true
	}
	if n, ok := a.([]complex128); ok {
		for _, el := range n {
			if f(el) == false {
				return false
			}
		}
		return true
	}
	return false
}
