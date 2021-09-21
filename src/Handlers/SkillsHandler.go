package Handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

type Skills struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	Icon    []string `json:"icon"`
}

var skills []Skills

func GetAllSkills(w http.ResponseWriter, _ *http.Request) {
	if os.Getenv("dev") == "true" {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	}

	getJsonData("skills", &skills)
	err := json.NewEncoder(w).Encode(skills)
	if err != nil {
		log.Fatal(err)
	}
}
