package handlers

import (
	"github.com/gorilla/mux"
	"gorm.io/gorm"
	db "gorm/db"
	"gorm/model"
	"net/http"
	"strconv"
)

func GetPredictions(rw http.ResponseWriter, r *http.Request) {

	info := model.PredictionsInfo{}
	db.Database.Find(&info)
	sendData(rw, info, http.StatusOK)

}
func GetPrediction(rw http.ResponseWriter, r *http.Request) {
	if info, err := getPredictionByDate(r); err != nil {
		sendError(rw, http.StatusNotFound)
	} else {
		sendData(rw, info, http.StatusOK)
	}

}

func getPredictionByDate(r *http.Request) (model.PredictionInfo, *gorm.DB) {
	vars := mux.Vars(r)
	var predictions model.PredictionInfo
	date, err := strconv.Atoi(vars["date"])
	if err != nil {
		http.Error(nil, "Fecha no válida", http.StatusBadRequest)
		return predictions, nil
	}
	db.Database.Where("date = ?", date).Find(&predictions)
	return predictions, nil

}

func CreatePrediction(name string, date int, predictionInfo string) {
	prediction := model.PredictionInfo{
		PlanetName:       name,
		Date:             date,
		WeatherCondition: predictionInfo,
	}
	db.Database.Create(&prediction)

}
