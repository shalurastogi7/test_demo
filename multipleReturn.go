package main

import "fmt"

func add(x int, y int) (int, int, int) {
	sum := x + y
	diff := x - y
	mul := x * y
	return sum, diff, mul

}

func main() {
	sum, diff, mul := add(5, 10)
	fmt.Print(sum, diff, mul)

}
