package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("Erica" == "Erica")
	fmt.Println("<", 10 < 8)
	fmt.Println(">", 10 > 8)
	fmt.Println("<=", 10 <= 8)
	fmt.Println(">=", 10 >= 8)
	fmt.Println("<=", 10 <= 10)
	fmt.Println(">=", 10 >= 10)

	d1 := time.Unix(0, 0)
	d2 := time.Unix(0, 0)

	fmt.Println("Data:", d1 == d2)
	fmt.Println("Data equal:", d1.Equal(d2))

	type Pessoa struct {
		Nome string
	}

	p1 := Pessoa{"Luiz"}
	p2 := Pessoa{"Luiz"}

	fmt.Println("Pessoa:", p1 == p2)

}
