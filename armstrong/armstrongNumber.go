package main

import "fmt"

func main() {

	var rightMost, num int
	var cubicSum = 0
	var tempNumber = 0

	fmt.Println("Enter a 3 digit number")
	fmt.Scanf("%d", &num)

	tempNumber = num

	for {

		rightMost = tempNumber % 10

		cubicSum += rightMost * rightMost * rightMost

		tempNumber /= 10

		if tempNumber == 0 {

			break
		}

	}

	if cubicSum == num {

		fmt.Println(num, " is an armstrong number")
	} else {

		fmt.Println(num, " is not an armstrong number")
	}

}
