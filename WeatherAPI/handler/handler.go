package handler

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

type Requests struct {
	Id            int    `json:"id"`
	City          string `json:"city"`
	TimeRequested string `json:"time_requested"`
	Temperature   string `json:"temperature"`
}

type Handlers struct {
}

func NewHandlers() *Handlers {
	return &Handlers{}
}

func (h *Handlers) Configure() *http.ServeMux {

	userMux := http.NewServeMux()
	userMux.HandleFunc("/weather", weatherAPI)
	userMux.HandleFunc("/listRequests", listRequests)

	return userMux
}

func listRequests(w http.ResponseWriter, r *http.Request) {
	var err error
	db, err = sql.Open("mysql", "root:1234@tcp(dbase:3306)/requestsdb") //127.0.0.1
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer db.Close()

	res, err := db.Query("SELECT * FROM requests")

	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer res.Close()
	requests := make([]Requests, 0)
	for res.Next() {

		var request Requests
		err := res.Scan(&request.Id, &request.City, &request.TimeRequested, &request.Temperature)

		if err != nil {
			log.Fatal(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		requests = append(requests, request)
	}
	requestsJson, err := json.Marshal(requests)

	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(requestsJson))
}

func weatherAPI(w http.ResponseWriter, r *http.Request) {

	cities, ok := r.URL.Query()["city"]

	if !ok || len(cities[0]) < 1 {
		log.Println("Url Param 'city' is missing")
		return
	}

	city := cities[0]

	data, err := query(city)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = db.Exec("INSERT INTO requests (city, time_requested, temperature) VALUES ($1, $2, $3)", city, time.Now().String(), fmt.Sprintf("%f", data.Main.Kelvin))

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(data)
}

func query(city string) (weatherData, error) {
	apiConfig, err := loadApiConfig(".apiConfig")
	if err != nil {
		return weatherData{}, err
	}

	resp, err := http.Get("http://api.openweathermap.org/data/2.5/weather?APPID=" + apiConfig.OpenWeatherMapApiKey + "&q=" + city)
	if err != nil {
		return weatherData{}, err
	}

	defer resp.Body.Close()

	var d weatherData

	if err := json.NewDecoder(resp.Body).Decode(&d); err != nil {
		return weatherData{}, err
	}

	return d, nil
}

func loadApiConfig(filename string) (apiConfigData, error) {
	bytes, err := ioutil.ReadFile(filename)

	if err != nil {
		return apiConfigData{}, err
	}

	var c apiConfigData
	err = json.Unmarshal(bytes, &c)
	if err != nil {
		return apiConfigData{}, err
	}

	return c, nil
}

type weatherData struct {
	Name string `json:"name"`
	Main struct {
		Kelvin float64 `json:"temp"`
	} `json:"main"`
}

type apiConfigData struct {
	OpenWeatherMapApiKey string `json:"OpenWeatherMapApiKey"`
}
