package main

import (
	"fmt"

	"github.com/h-hiwatashi/go-todolist/src/14/alib"
)

//test


func IsOne(n int) bool {
	if n == 1 {
		return true
	}
	return false
}

func main() {
	fmt.Println(IsOne(1)) // true
	fmt.Println(IsOne(2)) // false


	s:= []float64{1.0, 2.0, 3.0, 4.0, 5.0}
	fmt.Println(alib.Average(s)) // 3.0
}