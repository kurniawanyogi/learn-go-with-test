package main

import (
	"fmt"
	structs "learn-go-with-test/struct_method_interface/struct"
)

func main() {
	rectangle := structs.Rectangle{Width: 10.0, Height: 10.0}
	fmt.Println(Perimeter(rectangle))
	fmt.Println(Area(rectangle))
}

// this func created before refactor to method
func Perimeter(rectangle structs.Rectangle) float64 {
	return 2 * (rectangle.Height + rectangle.Width)
}
func Area(rectangle structs.Rectangle) float64 {
	return rectangle.Height * rectangle.Width
}
