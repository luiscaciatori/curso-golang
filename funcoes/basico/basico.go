package main

import (
	"fmt"
)

func F1() {
	fmt.Println("F1")
}

func F2(p1 string, p2 string) {
	fmt.Printf("F2: %s, %s\n", p1, p2)
}

func F3() string {
	return fmt.Sprint("F3\n")
}

func F4(p1, p2 string) string {
	return fmt.Sprintf("F4: %s, %s\n", p1, p2)
}

func F5(p1, p2 string) (string, string) {
	return p1, p2
}

func main() {
	F1()
	F2("Um", "Dois")

	_, r4 := F3(), F4("Um", "Dois")
	fmt.Printf("%s", r4)

	r51, r52 := F5("Dois", "Três")
	fmt.Printf("%s, %s", r51, r52)
}
