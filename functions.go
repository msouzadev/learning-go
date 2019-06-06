package main

import "fmt"

func otherFunc(a, b int) int {
	return a + b
}
func main() {
	myFunc := func(a, b int) int {
		return a + b
	}
	fmt.Println(myFunc(3, 5), otherFunc(6, 7))
}
