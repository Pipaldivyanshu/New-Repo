package main

import "fmt"

func incrementor() func() int {

	var x int

	return func() int {

		x++

		return x
	}
}
func main() {

	a := incrementor()
	b := incrementor()

	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(b())
	fmt.Println(b())
	fmt.Println(b())

}
