package every

import (
	"math"
	"math/cmplx"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInts(t *testing.T) {
	is := assert.New(t)

	oddInts := []int{-5, -3, -1, 1, 3, 5}
	evenInts := []int{-4, -2, 0, 2, 4}

	var isodd, iseven func(n int) bool
	isodd = func(n int) bool {
		return math.Abs(float64(n%2)) == 1
	}

	iseven = func(n int) bool {
		return math.Abs(float64(n%2)) == 0
	}

	is.True(Ints(oddInts, isodd))
	is.False(Ints(oddInts, iseven))
	is.True(Ints(evenInts, iseven))
	is.False(Ints(evenInts, isodd))

	oddUints := []int{1, 3, 5}
	evenUints := []int{0, 2, 4}

	isodd = func(n int) bool {
		return n%2 == 1
	}

	iseven = func(n int) bool {
		return n%2 == 0
	}

	is.True(Ints(oddUints, isodd))
	is.False(Ints(oddUints, iseven))
	is.True(Ints(evenUints, iseven))
	is.False(Ints(evenUints, isodd))
}

func TestFloats(t *testing.T) {
	is := assert.New(t)

	oddFloats := []float64{-5, -3, -1, 1, 3, 5}
	evenFloats := []float64{-4, -2, 0, 2, 4}

	var isodd, iseven func(n float64) bool
	isodd = func(n float64) bool {
		return math.Abs(float64(int(n)%2)) == 1
	}

	iseven = func(n float64) bool {
		return math.Abs(float64(int(n)%2)) == 0
	}

	is.True(Floats(oddFloats, isodd))
	is.False(Floats(oddFloats, iseven))
	is.True(Floats(evenFloats, iseven))
	is.False(Floats(evenFloats, isodd))

	oddUFloats := []float64{1, 3, 5}
	evenUFloats := []float64{0, 2, 4}

	isodd = func(n float64) bool {
		return int(n)%2 == 1
	}

	iseven = func(n float64) bool {
		return int(n)%2 == 0
	}

	is.True(Floats(oddUFloats, isodd))
	is.False(Floats(oddUFloats, iseven))
	is.True(Floats(evenUFloats, iseven))
	is.False(Floats(evenUFloats, isodd))
}

func TestStrings(t *testing.T) {
	is := assert.New(t)

	s := []string{"a b c", "a a", " ."}

	var fn func(s string) bool
	fn = func(s string) bool {
		r := []rune(s)
		for i := 0; i < len(r); i++ {
			if r[i] == ' ' {
				return true
			}
		}
		return false
	}

	is.True(Strings(s, fn))

	s = append(s, "abc")
	is.False(Strings(s, fn))

	s = []string{"apple", "banana"}
	fn = func(s string) bool {
		r := []rune(s)
		for i := 0; i < len(r); i++ {
			if r[i] == 'a' {
				return true
			}
		}
		return false
	}

	is.True(Strings(s, fn))
	s = append(s, "for you")
	is.False(Strings(s, fn))
}

func TestBytes(t *testing.T) {
	is := assert.New(t)

	s := []byte{'a', 'a'}

	var fn func(s byte) bool
	fn = func(s byte) bool {
		return s == 'a'
	}

	is.True(Bytes(s, fn))

	s = append(s, 'A')
	is.False(Bytes(s, fn))
}

func TestBools(t *testing.T) {
	is := assert.New(t)

	var b1, b2 []bool = []bool{true, true}, []bool{false, false}

	var isTrue, isFalse func(b bool) bool
	isTrue = func(b bool) bool {
		return b == true
	}
	isFalse = func(b bool) bool {
		return b == false
	}

	is.True(Bools(b1, isTrue))
	is.False(Bools(b2, isTrue))
	is.True(Bools(b2, isFalse))
	is.False(Bools(b1, isFalse))

	b1 = append(b1, b2...)
	is.False(Bools(b1, isTrue))
	is.False(Bools(b1, isFalse))
}

func TestComplexes(t *testing.T) {
	is := assert.New(t)

	c := []complex128{3 + 4i}
	var fn func(complex128) bool
	fn = func(c complex128) bool {
		return cmplx.Abs(c) == float64(5.0)
	}

	is.True(Complexes(c, fn))

	c = []complex128{cmplx.Inf()}
	is.True(Complexes(c, cmplx.IsInf))
}

func TestAny(t *testing.T) {
	is := assert.New(t)
	oddInts := []int{-5, -3, -1, 1, 3, 5}
	evenInts := []int{-4, -2, 0, 2, 4}

	var isodd, iseven func(n interface{}) bool
	isodd = func(num interface{}) bool {
		n, _ := num.(int)
		return math.Abs(float64(n%2)) == 1
	}

	iseven = func(num interface{}) bool {
		n, _ := num.(int)
		return math.Abs(float64(n%2)) == 0
	}

	is.True(Any(oddInts, isodd))
	is.False(Any(oddInts, iseven))
	is.True(Any(evenInts, iseven))
	is.False(Any(evenInts, isodd))

	oddUints := []int{1, 3, 5}
	evenUints := []int{0, 2, 4}

	isodd = func(num interface{}) bool {
		n, _ := num.(int)
		return n%2 == 1
	}

	iseven = func(num interface{}) bool {
		n, _ := num.(int)
		return n%2 == 0
	}

	is.True(Any(oddUints, isodd))
	is.False(Any(oddUints, iseven))
	is.True(Any(evenUints, iseven))
	is.False(Any(evenUints, isodd))

	oddFloats := []float64{-5, -3, -1, 1, 3, 5}
	evenFloats := []float64{-4, -2, 0, 2, 4}

	var isoddF, isevenF func(n interface{}) bool
	isoddF = func(num interface{}) bool {
		n, _ := num.(float64)
		return math.Abs(float64(int(n)%2)) == 1
	}

	isevenF = func(num interface{}) bool {
		n, _ := num.(float64)
		return math.Abs(float64(int(n)%2)) == 0
	}

	is.True(Any(oddFloats, isoddF))
	is.False(Any(oddFloats, isevenF))
	is.True(Any(evenFloats, isevenF))
	is.False(Any(evenFloats, isoddF))

	oddUFloats := []float64{1, 3, 5}
	evenUFloats := []float64{0, 2, 4}

	isoddF = func(num interface{}) bool {
		n, _ := num.(float64)
		return int(n)%2 == 1
	}

	isevenF = func(num interface{}) bool {
		n, _ := num.(float64)
		return int(n)%2 == 0
	}

	is.True(Any(oddUFloats, isoddF))
	is.False(Any(oddUFloats, isevenF))
	is.True(Any(evenUFloats, isevenF))
	is.False(Any(evenUFloats, isoddF))

	s := []string{"a b c", "a a", " ."}

	var fn func(s interface{}) bool
	fn = func(str interface{}) bool {
		s := str.(string)
		r := []rune(s)
		for i := 0; i < len(r); i++ {
			if r[i] == ' ' {
				return true
			}
		}
		return false
	}

	is.True(Any(s, fn))

	s = append(s, "abc")
	is.False(Any(s, fn))

	s = []string{"apple", "banana"}
	fn = func(str interface{}) bool {
		s := str.(string)
		r := []rune(s)
		for i := 0; i < len(r); i++ {
			if r[i] == 'a' {
				return true
			}
		}
		return false
	}

	is.True(Any(s, fn))
	s = append(s, "for you")
	is.False(Any(s, fn))

	b := []byte{'a', 'a'}

	var fnB func(b interface{}) bool
	fnB = func(by interface{}) bool {
		b := by.(byte)
		return b == 'a'
	}

	is.True(Any(b, fnB))

	b = append(b, 'A')
	is.False(Any(b, fnB))

	var b1, b2 []bool = []bool{true, true}, []bool{false, false}

	var isTrue, isFalse func(b interface{}) bool
	isTrue = func(b interface{}) bool {
		return b == true
	}
	isFalse = func(b interface{}) bool {
		return b == false
	}

	is.True(Any(b1, isTrue))
	is.False(Any(b2, isTrue))
	is.True(Any(b2, isFalse))
	is.False(Any(b1, isFalse))

	b1 = append(b1, b2...)
	is.False(Any(b1, isTrue))
	is.False(Any(b1, isFalse))

	c := []complex128{3 + 4i}
	var fnC func(interface{}) bool
	fnC = func(com interface{}) bool {
		c := com.(complex128)
		return cmplx.Abs(c) == float64(5.0)
	}

	is.True(Any(c, fnC))

	c = []complex128{cmplx.Inf()}
	is.True(Any(c, func(any interface{}) bool {
		a := any.(complex128)
		return cmplx.IsInf(a)
	}))

	var fi func(interface{}) bool
	fi = func(any interface{}) bool {
		a := any.(string)
		return a == "a"
	}

	var as interface{}
	as = []string{"a", "a"}

	is.True(Any(as, fi))
}
