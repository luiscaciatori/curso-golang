package main

import (
	"fmt"
)

func main() {
	array1 := [3]int{1, 2, 3}
	var slice1 []int

	fmt.Println(array1, slice1)

	slice2 := append(slice1, 4, 5, 6)
	fmt.Println(array1, slice2)

	slice3 := make([]int, 2)
	copy(slice3, slice2)

	fmt.Println(slice3)

	slice4 := make([]int, 3)
	copy(slice4, slice2)

	fmt.Println(slice4)
}
