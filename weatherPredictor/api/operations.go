package main

import "math"

func calculateCoordinates(planet Planet, day int) [2]float64 {
	angularVelocity := degreesToRadians(planet.AngularVelocity)
	radius := float64(planet.Radius)

	angle := angularVelocity * float64(day)
	x := radius * math.Cos(angle)
	y := radius * math.Sin(angle)

	return [2]float64{x, y}
}

func lcm(a, b int) int {
	return a * b / gcd(a, b)
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func degreesToRadians(degrees int) float64 {
	return float64(degrees) * (math.Pi / 180)
}
func isPointInsideTriangle(point, vertex1, vertex2, vertex3 [2]float64) bool {

	totalArea := calculateTriangleArea(vertex1, vertex2, vertex3)

	area1 := calculateTriangleArea(point, vertex1, vertex2)
	area2 := calculateTriangleArea(point, vertex2, vertex3)
	area3 := calculateTriangleArea(point, vertex3, vertex1)

	return math.Abs(totalArea-(area1+area2+area3)) <= 0.000001
}

func calculateTriangleArea(vertex1, vertex2, vertex3 [2]float64) float64 {
	a := calculateDistance(vertex1, vertex2)
	b := calculateDistance(vertex2, vertex3)
	c := calculateDistance(vertex3, vertex1)
	s := (a + b + c) / 2
	return math.Sqrt(s * (s - a) * (s - b) * (s - c))
}

func calculateTrianglePerimeter(vertex1, vertex2, vertex3 [2]float64) float64 {
	side1 := calculateDistance(vertex1, vertex2)
	side2 := calculateDistance(vertex2, vertex3)
	side3 := calculateDistance(vertex3, vertex1)

	perimeter := side1 + side2 + side3
	return perimeter
}

func calculateDistance(point1, point2 [2]float64) float64 {
	dx := point2[0] - point1[0]
	dy := point2[1] - point1[1]
	return math.Sqrt(dx*dx + dy*dy)
}
