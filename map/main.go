package main

import "fmt"

func main() {

	// var color map[string]string
	//color := make(map[int]string)

	//color[5] = "#ffff"

	color := map[string]string{

		"red":   "#ff0000",
		"green": "#4bf745",
		"white": "#ffff",
	}

	//delete(color, 5)
	//	fmt.Println(color)

	printMap(color)

}

func printMap(c map[string]string) {

	for color, hex := range c {

		fmt.Println("Hex code for  ", color, "is", hex)
	}
}
