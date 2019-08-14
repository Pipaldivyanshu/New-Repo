/*package main

import (
	"fmt"
	"math"
)

type vertex struct {
	x, y float64
}

func (v vertex) abs() float64 {

	return math.Sqrt(v.x*v.x + v.y*v.y)

}

func main() {

	v := vertex{3, 4}

	fmt.Println(v.abs())
}
*/

package main

import (
	"fmt"
	"math"
)

type vertex struct {
	X, Y float64
}

func (v *vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func (v vertex) abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := &vertex{3, 4}
	v.Scale(5)
	fmt.Println(v, v.abs())
}
