package main

import (
	"fmt"
	"reflect"
)

func main() {
	var a1 = [3]int{1, 2, 3}
	var s1 = []int{1, 2, 3}

	fmt.Println(a1, s1)

	var a2 = [5]int{1, 2, 3, 4, 5}
	var s2 = a1[1:3]

	fmt.Println(a2, s2)
	fmt.Println(reflect.TypeOf(a2), reflect.TypeOf(s2))

	var s3 = a2[:2]
	fmt.Println(s3)

	var s4 = s3[:1]
	fmt.Println(s4)

}
