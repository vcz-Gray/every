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
