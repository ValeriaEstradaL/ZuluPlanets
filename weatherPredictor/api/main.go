package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"gorm/handlers"
	"gorm/model"
	"log"
	"net/http"
)

func main() {
	model.MigrarInfo()
	planets := []Planet{
		{Name: "Ferengi", AngularVelocity: 1, Radius: 500},
		{Name: "Vulcano", AngularVelocity: -5, Radius: 1000},
		{Name: "Betazoide", AngularVelocity: 3, Radius: 2000},
	}
	daysInYear := 3650

	optimal, rain, maximus, drought := calculateOptimalConditionsPeriods(planets, daysInYear)

	fmt.Printf("Períodos de condiciones sequia  en los próximos 10 años: %d\n", drought)
	fmt.Printf("Períodos de lluvia en los próximos 10 años: %d\n", rain)
	fmt.Printf("Día del pico máximo de lluvia: %d\n", maximus)
	fmt.Printf("Períodos de condiciones óptimas de presión y temperatura: %d\n", optimal)

	mux := mux.NewRouter()

	mux.HandleFunc("/api/info/", handlers.GetPredictions).Methods("GET")
	mux.HandleFunc("/api/info/{date}", handlers.GetPrediction).Methods("GET")
	log.Fatal(http.ListenAndServe(":3000", mux))
}
