package model

import (
	"gorm/db"
)

type PredictionInfo struct {
	Id               int
	PlanetName       string
	Date             int
	WeatherCondition string
}

type PredictionsInfo []PredictionInfo

func MigrarInfo() {
	db.Database.AutoMigrate(PredictionInfo{})
}
