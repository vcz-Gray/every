package every

type ints func(int) bool
type floats func(float64) bool
type strings func(string) bool
type bytes func(byte) bool
type bools func(bool) bool
type complexes func(complex128) bool

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
