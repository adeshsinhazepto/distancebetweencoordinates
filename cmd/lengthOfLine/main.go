package main

import (
	"distancee/pkg/distance"
	"fmt"
)

func main() {
	var x1, y1, x2, y2 float64
	fmt.Println("Enter the coordinates of first point seperated by space:")
	fmt.Scanln(&x1, &y1)
	fmt.Println("Enter the coordinates of first point seperated by space:")
	fmt.Scanln(&x2, &y2)
	p1 := distance.NewPoint(int(x1), int(y1))
	p2 := distance.NewPoint(int(x2), int(y2))
	fmt.Println("Distance between the two points is:", distance.Distance(p1, p2))
}
