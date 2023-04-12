// main.go
package main

import (
	"fmt"
	. "water_cms/math"
	. "water_cms/projectA"
	. "water_cms/projectB"
)

func main() {
	fmt.Println("hello")
	fmt.Println(Average([]float64{1, 2}))

	fmt.Println(DoSomething())
	fmt.Println(DoSomethingElse())
}
