package every

import (
	"fmt"
	"math"
)

func ExampleInts() {
	fmt.Println(Ints([]int{-5, -3, -1, 1, 3, 5}, func(n int) bool {
		return int(math.Abs(float64(n)))%2 == 1
	}))
	fmt.Println(Ints([]int{-4, -2, 0, 2, 4}, func(n int) bool {
		return int(math.Abs(float64(n)))%2 == 0
	}))
}
