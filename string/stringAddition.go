package main

import "fmt"

func main() {

	var first string
	var second string

	fmt.Println("Enter first string")
	fmt.Scanln(&first)

	fmt.Println("Enter second string")
	fmt.Scanln(&second)

	fmt.Println(first + second)

}
