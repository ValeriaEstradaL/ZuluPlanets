package main

import (
	"gorm/handlers"
	"math"
)

type Planet struct {
	Name            string
	AngularVelocity int
	Radius          float64
}

func daysPlanetsAlignedWithNoSun(planets []Planet, days int) int {

	ferengiSpeed := planets[0].AngularVelocity
	vulcanoSpeed := planets[1].AngularVelocity
	betazoideSpeed := planets[2].AngularVelocity

	commonMultiple := lcm(lcm(int(math.Abs(float64(ferengiSpeed))), int(math.Abs(float64(vulcanoSpeed)))), int(math.Abs(float64(betazoideSpeed))))

	alignmentFrequency := days / commonMultiple

	return alignmentFrequency
}

func calculateOptimalConditionsPeriods(planets []Planet, daysInYear int) (int, int, int, int) {
	sun := [2]float64{0, 0}

	var ferengiCoords, vulcanoCoords, betazoideCoords [][2]float64
	rainyDays := 0
	maxRainDay := 0
	normalDays := 0
	droughtPeriods := 0

	for day := 1; day <= daysInYear; day++ {
		condition := ""
		ferengiCoord := calculateCoordinates(planets[0], day)
		vulcanoCoord := calculateCoordinates(planets[1], day)
		betazoideCoord := calculateCoordinates(planets[2], day)

		ferengiCoords = append(ferengiCoords, ferengiCoord)
		vulcanoCoords = append(vulcanoCoords, vulcanoCoord)
		betazoideCoords = append(betazoideCoords, betazoideCoord)

		if arePointsAlignedIncludeSun(ferengiCoord, vulcanoCoord, betazoideCoord, sun) {
			droughtPeriods++
			condition = "drought"

		} else if !arePointsAligned(ferengiCoord, vulcanoCoord, betazoideCoord) {
			if isPointInsideTriangle(sun, ferengiCoord, vulcanoCoord, betazoideCoord) {
				rainyDays++
				condition = "rainy"
			} else {
				normalDays++
				condition = "normal"
			}
		} else {
			condition = "Optimal"
		}
		if day <= 365 {
			handlers.CreatePrediction("Vulcano", day, condition)
		}
	}

	maxRainDay = findMaxRainDay(planets, daysInYear)

	optimalPeriods := daysPlanetsAlignedWithNoSun(planets, daysInYear)

	return optimalPeriods, rainyDays, maxRainDay, droughtPeriods
}

func arePointsAlignedIncludeSun(ferengi, vulcano, betazoide, sun [2]float64) bool {

	if ferengi[0] == sun[0] || vulcano[0] == sun[0] || betazoide[0] == sun[0] {
		return false
	}

	slope1 := (sun[1] - ferengi[1]) / (sun[0] - ferengi[0])
	slope2 := (sun[1] - vulcano[1]) / (sun[0] - vulcano[0])
	slope3 := (sun[1] - betazoide[1]) / (sun[0] - betazoide[0])

	epsilon := 1e-6
	if math.Abs(slope1-slope2) < epsilon && math.Abs(slope2-slope3) < epsilon {
		return true
	}

	return false
}

func arePointsAligned(point1, point2, point3 [2]float64) bool {
	slope1 := (point2[1] - point1[1]) / (point2[0] - point1[0])
	slope2 := (point3[1] - point2[1]) / (point3[0] - point2[0])

	return math.Abs(slope1-slope2) < 0.000001
}

func findMaxRainDay(planets []Planet, daysInYear int) int {
	var maxPerimeter float64
	var maxRainDay int

	for day := 1; day <= daysInYear; day++ {
		ferengiCoord := calculateCoordinates(planets[0], day)
		vulcanoCoord := calculateCoordinates(planets[1], day)
		betazoideCoord := calculateCoordinates(planets[2], day)
		sunCoord := [2]float64{0, 0}

		if isPointInsideTriangle(sunCoord, ferengiCoord, vulcanoCoord, betazoideCoord) {
			perimeter := calculateTrianglePerimeter(ferengiCoord, vulcanoCoord, betazoideCoord)
			if perimeter > maxPerimeter {
				maxPerimeter = perimeter
				maxRainDay = day
			}
		}
	}

	return maxRainDay
}
