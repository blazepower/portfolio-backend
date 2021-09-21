package Handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

type Experience struct {
	Header  string `json:"header"`
	Date    string `json:"date"`
	Content string `json:"content"`
}

var experience []Experience

func GetAllExperience(w http.ResponseWriter, _ *http.Request) {
	if os.Getenv("dev") == "true" {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	}

	getJsonData("experience", &experience)
	err := json.NewEncoder(w).Encode(experience)
	if err != nil {
		log.Fatal(err)
	}
}

func GetAllEducation(w http.ResponseWriter, _ *http.Request)  {
	if os.Getenv("dev") == "true" {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	}

	getJsonData("education", &experience)
	err := json.NewEncoder(w).Encode(experience)
	if err != nil {
		log.Fatal(err)
	}
}