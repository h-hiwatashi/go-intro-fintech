package main

import (
	"fmt"
	"strconv"
)

func PrintSlice[T any](s []T) {
	for _, v := range s {
		println(v)
	}
}

func f[T fmt.Stringer](xs []T) []string {
	result := []string{}
	for _, x := range xs {
		result = append(result, x.String())
	}
	return result
}

type MyInt int

func (i MyInt) String() string {
	return strconv.Itoa(int(i))
}

func main() {
	PrintSlice([]int{1, 2, 3})
	PrintSlice([]string{"a", "b", "c"})
	PrintSlice([]float64{1.1, 2.2, 3.3})
	fmt.Println(f([]MyInt{1, 2, 3}))
}
