package main

import "fmt"

func main() {
	var l, u, x, i int
	fmt.Println("enter the 1 m")
	fmt.Scanln(&l)
	fmt.Println("enter the 2 no")
	fmt.Scanln(&u)

	for x = l + 1; x <= u-1; x++ {
		for i = 2; i <= x; i++ {
			if x%i == 0 {
				break
			}
		}
		if x == i {
			fmt.Println(x)

		}

	}

}
