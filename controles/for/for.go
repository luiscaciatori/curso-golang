package main

import (
	"fmt"
	"time"
)

func main() {
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i++
	}

	for j := 0; j <= 20; j += 2 {
		fmt.Printf("%d ", j)
	}

	fmt.Println("")

	for k := 0; k <= 10; k++ {
		if k%2 == 0 {
			fmt.Printf("Par ")
		} else {
			fmt.Printf("Impar ")
		}
	}

	for {
		fmt.Println("Esse é infinito")
		time.Sleep(time.Second * 5)
	}
}
