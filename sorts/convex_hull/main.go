package main

import (
	"fmt"
	"math"
	"sort"
)

type Point struct {
	x, y int
}

// A function to find the point with the lowest y-coordinate
func lowestPoint(points []Point) Point {
	lp := points[0]
	for _, p := range points {
		if p.y < lp.y || (p.y == lp.y && p.x < lp.x) {
			lp = p
		}
	}
	return lp
}

// A function to calculate the polar angle between two points
func polarAngle(p1, p2 Point) float64 {
	return math.Atan2(float64(p2.y-p1.y), float64(p2.x-p1.x))
}

// A function to calculate the cross product of three points
func crossProduct(p0, p1, p2 Point) int {
	return (p1.x-p0.x)*(p2.y-p0.y) - (p1.y-p0.y)*(p2.x-p0.x)
}

// Graham's scan algorithm to find the convex hull
func convexHull(points []Point) []Point {
	if len(points) < 3 {
		return points
	}

	lp := lowestPoint(points)

	sort.Slice(points, func(i, j int) bool {
		angle1 := polarAngle(lp, points[i])
		angle2 := polarAngle(lp, points[j])
		if angle1 == angle2 {
			return distance(lp, points[i]) < distance(lp, points[j])
		}
		return angle1 < angle2
	})

	hull := []Point{lp}
	for _, p := range points {
		for len(hull) > 1 && crossProduct(hull[len(hull)-2], hull[len(hull)-1], p) <= 0 {
			hull = hull[:len(hull)-1]
		}
		hull = append(hull, p)
	}

	return hull
}

// A function to calculate the distance between two points
func distance(p1, p2 Point) float64 {
	return math.Sqrt(float64((p1.x-p2.x)*(p1.x-p2.x) + (p1.y-p2.y)*(p1.y-p2.y)))
}

func main() {
	points := []Point{
		{0, 3}, {1, 1}, {2, 2}, {4, 4},
		{0, 0}, {1, 2}, {3, 1}, {3, 3},
	}
	hull := convexHull(points)
	fmt.Println("Convex Hull:")
	for _, p := range hull {
		fmt.Printf("(%d, %d)\n", p.x, p.y)
	}
}
