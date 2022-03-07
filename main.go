package main

import (
	"cwb-sdk-go/model"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	KEY := os.Getenv("API_KEY")
	client, _ := New(KEY)
	w, _ := client.GeneralWeatherForecast36()
	log.Print(w.Locations[0].CI)
}

func New(API_KEY string) (*model.Client, error) {
	return &model.Client{
		API_KEY:  API_KEY,
		BASE_URL: "https://opendata.cwb.gov.tw/api/v1/rest/datastore/%s?Authorization=%s"}, nil
}
