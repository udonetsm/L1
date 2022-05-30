package main

import (
	"L1/point"
	"log"
)

func compute(x1, y1, x2, y2 float64) {
	// make two new obect
	point1 := point.NewPoint()
	point2 := point.NewPoint()
	// set values for both of incapsulated params
	point1.SetParams(x1, y1)
	point2.SetParams(x2, y2)
	// compute distance for current params
	log.Println(point.Res(point1, point2))
}
