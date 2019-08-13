package main

import "fmt"

func main() {

	var num, remainder, temp int
	var reverse = 0

	fmt.Println("Enter the number")
	fmt.Scanf("%d", &num)

	temp = num

	for {

		remainder = num % 10

		reverse = reverse*10 + remainder

		num = num / 10

		if num == 0 {

			break
		}

	}

	if temp == reverse {

		fmt.Println(temp, "is a palindrome")
	} else {

		fmt.Println(temp, "is not a palindrome")
	}

}
