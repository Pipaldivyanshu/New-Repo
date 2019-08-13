package main

import "fmt"

func main() {

	var s int
	var len int
	var br int

	fmt.Println("Enter the value of s")
	fmt.Scanf("%d", &s)

	fmt.Println("Enter the value of length and breadth")
	fmt.Scanf("%d", &len)
	fmt.Scanf("%d", &br)

	var rectArea int
	var sqArea int

	rectArea = len * br
	sqArea = s * s

	fmt.Println(rectArea, "is the area of the rectangle")
	fmt.Println(sqArea, "is the area of the square")
}
