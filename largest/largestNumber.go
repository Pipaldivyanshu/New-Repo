package main

import "fmt"

func main() {

	var first int
	var second int
	var third int

	fmt.Println("Enter the first number")
	fmt.Scanf("%d", &first)

	fmt.Println("Enter the second number")
	fmt.Scanf("%d", &second)

	fmt.Println("Enter the third number")
	fmt.Scanf("%d", &third)

	if first > second && first > third {

		fmt.Println("First number is the largest")

	} else if second > first && second > third {

		fmt.Println("Second number is the largest")

	} else {

		fmt.Println("Third number is the largest")
	}

}
