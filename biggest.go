package main

import "fmt"

func big(x int, y int, z int) {
	if x < y {
		fmt.Println("y is greater ", y)
	} else if y < z {
		fmt.Println("z is greater ", z)

	} else {
		fmt.Println("x is greater ", x)

	}

}

func main() {
	big(5, 10, 8)

}
