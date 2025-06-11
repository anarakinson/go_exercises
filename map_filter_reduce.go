package main

import (
	"fmt"
)

func main() {

	// map example
	var lst = []int{1, 2, 3, 4, 5, 6}

	lst = Map(
		lst, func(x int) int {
			return x * x
		},
	)

	for _, v := range lst {
		fmt.Print(v, " ")
	}
	fmt.Println()
	// map example end
	//--------------------
	// filter example
	lst = Filter(
		lst, func(x int) bool {
			return x%2 == 0
		},
	)

	for _, v := range lst {
		fmt.Print(v, " ")
	}
	fmt.Println()
	// filter example end
	//--------------------
	// reduce example
	reduced := Reduce(
		lst,
		func(x int, y int) int {
			return x * y
		},
		1,
	)

	fmt.Println(reduced)
	// reduce example end

}

// define numeric type
type Numeric_t interface {
	int | int8 | int16 | int32 | int64 |
		uint | uint8 | uint16 | uint32 | uint64 |
		float32 | float64
}

// function applied to every member of array
func Map[T Numeric_t](lst []T, applyer func(T) T) []T {
	res := []T{}
	for _, v := range lst {
		res = append(res, applyer(v))
	}
	return res
}

// filtering every member of array
func Filter[T Numeric_t](lst []T, condition func(T) bool) []T {
	res := []T{}
	for _, v := range lst {
		if condition(v) {
			res = append(res, v)
		}
	}
	return res
}

// reduce all members of array into one
func Reduce[T Numeric_t](lst []T, reduce_func func(T, T) T, init_val T) T {
	for i := 0; i < len(lst); i++ {
		val := lst[i]
		init_val = reduce_func(val, init_val)
	}
	return init_val
}
