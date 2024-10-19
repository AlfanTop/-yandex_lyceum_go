package main

import (
	"fmt"
	"math"
)

func SqRoots1() {
	var a, b, c float64
	fmt.Scanln(&a, &b, &c)

	D := b*b - 4*a*c

	if D > 0 {
		root1 := (-b + math.Sqrt(D)) / (2 * a)
		root2 := (-b - math.Sqrt(D)) / (2 * a)
		if root1 < root2 {
			fmt.Printf("%.6f %.6f\\n", root1, root2)
		} else {
			fmt.Printf("%.6f %.6f\\n", root2, root1)
		}
	} else if D == 0 {
		root := -b / (2 * a)
		fmt.Printf("%.6f %.6f\\n", root, root)
	} else {
		fmt.Println("0 0")
	}
}
