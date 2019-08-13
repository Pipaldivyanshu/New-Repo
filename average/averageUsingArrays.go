package main

import "fmt"

func main() {

	var arr [100]int
	var avg, num, sum int

	fmt.Println("Enter the number of elements ")
	fmt.Scanln(&num)

	for i := 0; i < num; i++ {

		fmt.Println("Enter the numbers")
		fmt.Scanln(&arr[i])

		sum = sum + arr[i]

	}

	avg = sum / num

	fmt.Println("Average is", avg)

}
