package main

import "fmt"

func sayHello(msg string) {

	fmt.Println(msg)
}

func main() {

	sayHello("Hello")

	//anonymous function declared
	func() {

		fmt.Println("Anonymous function")
	}()

	func(x int) {

		fmt.Println("I am going to be ", x)
	}(27)

}
