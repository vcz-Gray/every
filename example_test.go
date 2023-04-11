package every

import (
	"fmt"
	"math"
	"math/cmplx"
	strs "strings"
)

func ExampleInts() {
	fmt.Println(Ints([]int{-5, -3, -1, 1, 3, 5}, func(n int) bool {
		return int(math.Abs(float64(n)))%2 == 1
	}))
	fmt.Println(Ints([]int{-4, -2, 0, 2, 4}, func(n int) bool {
		return int(math.Abs(float64(n)))%2 == 0
	}))
	// Output:
	// true
	// true
}

func ExampleFloats() {
	fmt.Println(Floats([]float64{10.1, 10.2, 10.3, 10.4, 10.5}, func(n float64) bool {
		return math.Ceil(n) == float64(11)
	}))
	// Output:
	// true
}

func ExampleStrings() {
	fmt.Println(Strings([]string{"An apple", "A banana", "A guaba"}, func(s string) bool {
		return strs.HasPrefix(s, "An")
	}))

	fmt.Println(Strings([]string{"+82-10-0000-0000", "+82-10-4654-0000"}, func(s string) bool {
		return strs.HasPrefix(s, "+82-10")
	}))
	// Output:
	// false
	// true
}

func ExampleBytes() {
	fmt.Println(Bytes([]byte{'A', 'B', 'C', 'D'}, func(b byte) bool {
		return b < 'a'
	}))
	// Output:
	// true
}

func ExampleBools() {
	fmt.Println(Bools([]bool{true, true, true}, func(b bool) bool {
		return b == true
	}))
	// Output:
	// true
}

func ExampleComplexes() {
	fmt.Println(Complexes([]complex128{cmplx.Inf()}, func(c complex128) bool {
		return cmplx.IsInf(c)
	}))
	// Output:
	// true
}
