package main

import (
	"fmt"
	"math"
)

func SqRoots() {
	var a, b, c float64
	fmt.Scanln(&a, &b, &c)

	d := b*b - 4*a*c
	if d > 0 {
		root1 := (-b + math.Sqrt(d)) / (2 * a)
		root2 := (-b - math.Sqrt(d)) / (2 * a)
		if root1 < root2 {
			fmt.Printf("%.3f %.3f\\n", root1, root2)
		} else {
			fmt.Printf("%.3f %.3f\\n", root2, root1)
		}
	} else if d == 0 {
		root := -b / (2 * a)
		fmt.Printf("%.3f\\n", root)
	} else {
		fmt.Println("0 0")
	}
}
func main() {
	SqRoots()
}
