package main

import "fmt"

func digit() {
	var num, i, count int
	count = 0
	fmt.Scanln(&num)
	for i = 1; num != 0; i++ {
		num = num / 10
		count++

	}
	fmt.Println(count)

}

func main() {
	digit()

}
