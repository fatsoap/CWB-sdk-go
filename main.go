package main

import (
	"cwb-sdk-go/model"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

var BASE_URL = "https://opendata.cwb.gov.tw/api"

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	res, err := http.Get(BASE_URL + "/v1/rest/datastore/F-D0047-001?Authorization=" + os.Getenv("API_KEY"))
	if err != nil {
		log.Fatal("Error When Fetching Data")
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal("Error When Reading Body")
	}
	// sb := string(body)
	var A model.TaiwanForecast36
	json.Unmarshal(body, &A)
	log.Print(A.Records.Location[0].WeatherElements[0].Times[0].Parameter.ParameterValue)
}
