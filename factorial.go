package main

import "fmt"

func main() {
	var fact int
	fact = 1
	var num int
	fmt.Println("enter no")

	fmt.Scanln(&num)
	var i int
	for i = 1; i <= num; i++ {
		fact = fact * i
	}
	fmt.Println(fact)
}
