package main

import "fmt"

func factor() {
	var i, num int
	fmt.Println("enter the no")
	fmt.Scanln(&num)
	for i = 1; i <= num; i++ {
		if num%i == 0 {
			fmt.Print(" ", i)
		}
	}
}

func main() {
	factor()
}
